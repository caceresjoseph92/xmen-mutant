package xmen

import (
	"xmen-mutant/kit/event"
)

const PersonCreatedEventType event.Type = "events.person.created"

type PersonCreatedEvent struct {
	event.BaseEvent
	mutant bool
	dna    []string
}

func NewPersonCreatedEvent(mutant bool, dna []string) PersonCreatedEvent {
	return PersonCreatedEvent{
		mutant: mutant,
		dna:    dna,

		BaseEvent: event.NewBaseEvent(mutant, dna),
	}
}

func (e PersonCreatedEvent) Type() event.Type {
	return PersonCreatedEventType
}

func (e PersonCreatedEvent) PersonMutant() bool {
	return e.mutant
}

func (e PersonCreatedEvent) PersonDna() []string {
	return e.dna
}
