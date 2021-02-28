package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Bus defines the expected behaviour from an event bus.
type Bus interface {
	// Publish is the method used to publish new events.
	Publish(context.Context, []Event) error
	// Subscribe is the method used to subscribe new event handlers.
	Subscribe(Type, Handler)
}

//mockery --case=snake --outpkg=eventmocks --output=eventmocks --name=Bus

// Handler defines the expected behaviour from an event handler.
type Handler interface {
	Handle(context.Context, Event) error
}

// Type represents a domain event type.
type Type string

// Event represents a domain command.
type Event interface {
	ID() string
	AggregateDNS() []string
	OccurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	eventID         string
	aggregateMutant bool
	aggregateDNS    []string
	occurredOn      time.Time
}

func NewBaseEvent(aggregateMutant bool, aggregateDNS []string) BaseEvent {
	return BaseEvent{
		eventID:         uuid.New().String(),
		aggregateMutant: aggregateMutant,
		aggregateDNS:    aggregateDNS,
		occurredOn:      time.Now(),
	}
}

func (b BaseEvent) ID() string {
	return b.eventID
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateDNS() []string {
	return b.aggregateDNS
}
