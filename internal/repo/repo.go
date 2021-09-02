package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-meeting-api/internal/models"
)

// MeetingRepo storage interface of Meeting Entity
type MeetingRepo interface {
	AddMeetings(ctx context.Context, meetings []models.Meeting) error
	ListMeetings(limit, offset uint64) ([]models.Meeting, error)
	DescribeMeeting(meetingId uuid.UUID) (models.Meeting, error)
	DeleteMeeting(ctx context.Context, meetingId uuid.UUID) error
	UpdateMeeting(ctx context.Context, meeting models.Meeting) error
}

type repo struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewRepo(db *sqlx.DB) MeetingRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) DescribeMeeting(meetingId uuid.UUID) (models.Meeting, error) {
	var meeting models.Meeting
	var userString string
	query, args, err := squirrel.
		Select("m.id, m.title, m.creator, m.meeting_date, ms.id as state_id, ms.name as state_name, mu.users").
		From("meetings as m").
		InnerJoin("meeting_states as ms on m.state_id = ms.id").
		LeftJoin("(select meeting_id, json_agg(user_id) as users from meeting_users group by meeting_id) as mu on m.id = mu.meeting_id").
		Where(squirrel.Eq{"m.id": meetingId}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return meeting, err
	}

	row := r.db.QueryRowx(query, args...)

	err = row.Scan(&meeting.ID, &meeting.Title, &meeting.UserID, &meeting.Date, &meeting.State.ID, &meeting.State.Name, &userString)
	if err == sql.ErrNoRows {
		return meeting, errors.New("Meeting with id " + meetingId.String() + " not found")
	} else if err != nil {
		return meeting, err
	}
	err = json.Unmarshal([]byte(userString), &meeting.Users)
	if err != nil {
		return meeting, err
	}
	return meeting, nil
}

func (r *repo) addMeeting(ctx context.Context, meeting models.Meeting) error {
	meeting.ID = uuid.New()
	err := r.checkMeetingState(&meeting.State)
	if err != nil {
		return err
	}
	query, args, err := squirrel.
		Insert("meetings").
		Columns("id", "creator", "title", "state_id", "meeting_date").
		Values(meeting.ID, meeting.UserID, meeting.Title, meeting.State.ID, meeting.Date).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}

	for _, meetingUser := range meeting.Users {
		addMeetingUser, args, err := squirrel.
			Insert("meeting_users").
			Columns("meeting_id", "user_id").
			Values(meeting.ID, meetingUser).
			PlaceholderFormat(squirrel.Dollar).
			ToSql()
		if err != nil {
			errTx := tx.Rollback()
			if errTx != nil {
				return errTx
			}
			return err
		}

		_, err = tx.ExecContext(ctx, addMeetingUser, args...)
		if err != nil {
			errTx := tx.Rollback()
			if errTx != nil {
				return errTx
			}
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) AddMeetings(ctx context.Context, meetings []models.Meeting) error {
	for _, meeting := range meetings {
		err := r.addMeeting(ctx, meeting)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *repo) ListMeetings(limit, offset uint64) ([]models.Meeting, error) {
	var meeting models.Meeting
	result := make([]models.Meeting, 0, limit)

	query, args, err := squirrel.
		Select("m.id, m.title, m.creator, m.meeting_date, ms.id as state_id, ms.name as state_name, mu.users").
		From("meetings as m").
		InnerJoin("meeting_states as ms on m.state_id = ms.id").
		LeftJoin("(select meeting_id, json_agg(user_id) as users from meeting_users group by meeting_id) as mu on m.id = mu.meeting_id").
		OrderBy("m.meeting_date").
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err = rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	for rows.Next() {
		var userString string
		if err = rows.Scan(&meeting.ID, &meeting.Title, &meeting.UserID, &meeting.Date, &meeting.State.ID, &meeting.State.Name, &userString); err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(userString), &meeting.Users)
		if err != nil {
			return nil, err
		}

		result = append(result, meeting)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) UpdateMeeting(ctx context.Context, meeting models.Meeting) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = r.checkMeetingState(&meeting.State)
	if err != nil {
		return err
	}
	meetingsQuery, args, err := squirrel.
		Update("meetings").
		Set("creator", meeting.UserID).
		Set("title", meeting.Title).
		Set("state_id", meeting.State.ID).
		Set("meeting_date", meeting.Date).
		Where(squirrel.Eq{"id": meeting.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	clearMeetingUsers, args, err := squirrel.
		Delete("meeting_users").
		Where(squirrel.Eq{"meeting_id": meeting.ID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, meetingsQuery, args...)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}

	_, err = tx.ExecContext(ctx, clearMeetingUsers, args...)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}
	for _, meetingUser := range meeting.Users {
		addMeetingUser, args, err := squirrel.
			Insert("meeting_users").
			Columns("meeting_id", "user_id").
			Values(meeting.ID, meetingUser).
			PlaceholderFormat(squirrel.Dollar).
			ToSql()
		if err != nil {
			errTx := tx.Rollback()
			if errTx != nil {
				return errTx
			}
			return err
		}

		_, err = tx.ExecContext(ctx, addMeetingUser, args...)
		if err != nil {
			errTx := tx.Rollback()
			if errTx != nil {
				return errTx
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteMeeting(ctx context.Context, meetingId uuid.UUID) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	queryMeetings, args, err := squirrel.
		Delete("meetings").
		Where(squirrel.Eq{"id": meetingId}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	queryMeetingUsers, args, err := squirrel.
		Delete("meeting_users").
		Where(squirrel.Eq{"meeting_id": meetingId}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryMeetings, args...)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}
	_, err = tx.ExecContext(ctx, queryMeetingUsers, args...)
	if err != nil {
		errTx := tx.Rollback()
		if errTx != nil {
			return errTx
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) checkMeetingState(state *models.MeetingState) error {
	var resState models.MeetingState
	builder := squirrel.
		Select("id as state_id, name as state_name").
		From("meeting_states").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"name": state.Name})
	if state.ID != 0 {
		builder.Where(squirrel.Eq{"id": state.ID})
	}
	builder.Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	row := r.db.QueryRowx(query, args...)
	err = row.StructScan(&resState)
	if err == sql.ErrNoRows {
		insertQuery, args, err := squirrel.
			Insert("meeting_states").
			Columns("name").
			Values(state.Name).
			PlaceholderFormat(squirrel.Dollar).
			Suffix("RETURNING \"id\"").
			ToSql()
		if err != nil {
			log.Println(1)
			log.Println(err)
		}
		row := r.db.QueryRowx(insertQuery, args...)
		if err != nil {
			return err
		}
		err = row.Scan(&resState.ID)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	state.ID = resState.ID
	return nil
}
