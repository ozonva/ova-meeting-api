package producer

import (
	"encoding/json"
	"github.com/google/uuid"

	"github.com/Shopify/sarama"
)

type producer struct {
	prod sarama.SyncProducer
}

type event struct {
	Type string    `json:"type"`
	Id   uuid.UUID `json:"id"`
}

func New(brokers []string) (Producer, error) {
	conf := sarama.NewConfig()
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	prod, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return nil, err
	}

	return &producer{
		prod: prod,
	}, nil
}

func (p *producer) Send(topic string, event Event) error {
	prodMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(event.Value()),
	}

	_, _, err := p.prod.SendMessage(prodMsg)
	return err
}

func (p *producer) Close() {
	err := p.prod.Close()
	if err != nil {
		return
	}
}

func PrepareEvent(typ EventType, id uuid.UUID) Event {
	return &event{
		Type: typ.String(),
		Id:   id,
	}
}

func (e *event) Value() string {
	if e == nil {
		return ""
	}

	msg, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(msg)
}
