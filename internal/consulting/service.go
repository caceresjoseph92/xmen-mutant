package consulting

import (
	"context"

	xmen "xmen-mutant/internal"
	"xmen-mutant/kit/event"
)

// PersonService is the default PersonService interface
// implementation returned by consulting.NewPersonService.
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

// ConsultPerson implements the consulting.PersonService interface.
func (s PersonService) ConsultPerson(ctx context.Context, args map[string]interface{}) (stats map[string]interface{}, err error) {
	response, err := s.personRepository.Consult(ctx, args)
	if err != nil {
		return nil, err
	}
	s.eventBus.Publish(ctx, nil)

	return response, nil
}
