package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	create,
	update,
	delete prometheus.Counter
}

func New() Metrics {
	return &metrics{
		create: promauto.NewCounter(prometheus.CounterOpts{
			Name: "meeting_create_successful",
			Help: "Number of successful created events",
		}),
		update: promauto.NewCounter(prometheus.CounterOpts{
			Name: "meeting_update_successful",
			Help: "Number of successful updated events",
		}),
		delete: promauto.NewCounter(prometheus.CounterOpts{
			Name: "meeting_delete_successful",
			Help: "Number of successful deleted events",
		}),
	}
}

func (m *metrics) IncCreate() {
	m.create.Inc()
}

func (m *metrics) IncUpdate() {
	m.update.Inc()
}

func (m *metrics) IncDelete() {
	m.delete.Inc()
}
