package creating

import (
	"context"

	xmen "xmen-mutant/internal"
	"xmen-mutant/kit/event"
)

// PersonService is the default PersonService interface
// implementation returned by creating.NewPersonService.
type PersonService struct {
	personRepository xmen.PersonRepository
	eventBus         event.Bus
}

// NewPersonService returns the default Service interface implementation.
func NewPersonService(personRepository xmen.PersonRepository, eventBus event.Bus) PersonService {
	return PersonService{
		personRepository: personRepository,
		eventBus:         eventBus,
	}
}

// CreatePerson implements the creating.PersonService interface.
func (s PersonService) CreatePerson(ctx context.Context, mutant bool, dna []string) error {
	person, err := xmen.NewPerson(mutant, dna)
	if err != nil {
		return err
	}

	if err := s.personRepository.Save(ctx, person); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, person.PullEvents())
}
