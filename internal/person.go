package xmen

import (
	"context"
	"errors"

	"xmen-mutant/kit/event"
)

// PersonMutant represents the person mutant.
type PersonMutant struct {
	value bool
}

// Bool type converts the PersonMutant into bool.
func (mutant PersonMutant) Bool() bool {
	return mutant.value
}

var ErrEmptyDna = errors.New("the field Dna can not be empty")

// PersonDna represents the person dna.
type PersonDna struct {
	value []string
}

func NewMutant(value bool) (PersonMutant, error) {
	return PersonMutant{
		value: value,
	}, nil
}

func NewPersonDna(value []string) (PersonDna, error) {
	if value == nil {
		return PersonDna{}, ErrEmptyDna
	}

	return PersonDna{
		value: value,
	}, nil
}

// String type converts the PersonDuration into string.
func (dna PersonDna) String() []string {
	return dna.value
}

// Person is the data structure that represents a person.
type Person struct {
	mutant PersonMutant
	dna    PersonDna

	events []event.Event
}

// PersonRepository defines the expected behaviour from a person storage.
type PersonRepository interface {
	Save(ctx context.Context, person Person) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=PersonRepository

// NewPerson creates a new person.
func NewPerson(mutant bool, dna []string) (Person, error) {
	dnaVO, err := NewPersonDna(dna)
	mutantVO, err := NewMutant(mutant)
	if err != nil {
		return Person{}, err
	}

	person := Person{
		mutant: mutantVO,
		dna:    dnaVO,
	}
	person.Record(NewPersonCreatedEvent(mutant, dnaVO.String()))
	return person, nil
}

// Mutant returns the mutant name.
func (c Person) Mutant() PersonMutant {
	return c.mutant
}

// Dna returns the person dna.
func (c Person) Dna() PersonDna {
	return c.dna
}

// Record records a new domain event.
func (c *Person) Record(evt event.Event) {
	c.events = append(c.events, evt)
}

// PullEvents returns all the recorded domain events.
func (c Person) PullEvents() []event.Event {
	evt := c.events
	c.events = []event.Event{}

	return evt
}
