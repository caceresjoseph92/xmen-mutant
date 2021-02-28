package consulting

import (
	"context"
	"errors"
	"testing"
	xmen "xmen-mutant/internal"
	"xmen-mutant/internal/platform/storage/storagemocks"
	"xmen-mutant/kit/event"
	"xmen-mutant/kit/event/eventmocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PersonService_ConsultPerson_RepositoryError(t *testing.T) {
	arg := map[string]interface{}{}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Consult", mock.Anything, mock.Anything).Return(nil, errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	_, err := personService.ConsultPerson(context.Background(), arg)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_PersonService_ConsultPerson_EventsBusError(t *testing.T) {
	arg := map[string]interface{}{}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Consult", mock.Anything, mock.Anything).Return(nil, nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	_, err := personService.ConsultPerson(context.Background(), arg)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_PersonService_ConsultPerson_Succeed(t *testing.T) {
	arg := map[string]interface{}{}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Consult", mock.Anything, mock.AnythingOfType("xmen.Person")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
		evt := events[0].(xmen.PersonCreatedEvent)
		return evt.PersonMutant() == true
	})).Return(nil)

	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	_, err := personService.ConsultPerson(context.Background(), arg)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
