package creating

import (
	"context"
	"errors"
	xmen "xmen-mutant/internal"
	"xmen-mutant/internal/increasing"
	"xmen-mutant/kit/event"
)

type IncreasePersonsCounterOnPersonCreated struct {
	increasingService increasing.PersonCounterService
}

func NewIncreasePersonsCounterOnPersonCreated(increaserService increasing.PersonCounterService) IncreasePersonsCounterOnPersonCreated {
	return IncreasePersonsCounterOnPersonCreated{
		increasingService: increaserService,
	}
}

func (e IncreasePersonsCounterOnPersonCreated) Handle(_ context.Context, evt event.Event) error {
	personCreatedEvt, ok := evt.(xmen.PersonCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(personCreatedEvt.ID())
}
