package creating

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

func Test_PersonService_CreatePerson_RepositoryError(t *testing.T) {
	personMutant := true
	personDna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("xmen.Person")).Return(errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	err := personService.CreatePerson(context.Background(), personMutant, personDna)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_PersonService_CreatePerson_EventsBusError(t *testing.T) {
	personMutant := true
	personDna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("xmen.Person")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	err := personService.CreatePerson(context.Background(), personMutant, personDna)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_PersonService_CreatePerson_Succeed(t *testing.T) {
	personMutant := true
	personDna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("xmen.Person")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
		evt := events[0].(xmen.PersonCreatedEvent)
		return evt.PersonMutant() == personMutant
	})).Return(nil)

	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	err := personService.CreatePerson(context.Background(), personMutant, personDna)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
