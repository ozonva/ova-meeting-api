-- +goose Up
CREATE TABLE meeting_states
(
    id   serial  NOT NULL,
    name varchar NOT NULL,
    CONSTRAINT meeting_states_pk PRIMARY KEY (id)
);

CREATE TABLE meetings
(
    id           uuid      NOT NULL,
    title        varchar   NULL,
    creator      int       NOT NULL,
    meeting_date timestamp NULL,
    state_id     int       NOT NULL,
    CONSTRAINT meetings_pk PRIMARY KEY (id),
    CONSTRAINT meetings_fk FOREIGN KEY (state_id) REFERENCES meeting_states (id)
);

CREATE TABLE meeting_users
(
    meeting_id uuid NULL,
    user_id    int  NULL,
    CONSTRAINT meeting_users_fk FOREIGN KEY (meeting_id) REFERENCES meetings (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down