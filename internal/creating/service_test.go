package creating

import (
	"context"
	"errors"
	"testing"

	"xmen-mutant/internal/platform/storage/storagemocks"
	"xmen-mutant/kit/event/eventmocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PersonService_CreatePerson_RepositoryError(t *testing.T) {
	personMutant := true
	personDna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("xmen.Person")).Return(nil, errors.New("something unexpected happened"))

	eventBusMock := new(eventmocks.Bus)

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	_, err := personService.CreatePerson(context.Background(), personMutant, personDna)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_PersonService_CreatePerson_EventsBusSucceed(t *testing.T) {
	personMutant := true
	personDna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	personRepositoryMock := new(storagemocks.PersonRepository)
	personRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("xmen.Person")).Return(nil, nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	personService := NewPersonService(personRepositoryMock, eventBusMock)

	_, err := personService.CreatePerson(context.Background(), personMutant, personDna)

	personRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
