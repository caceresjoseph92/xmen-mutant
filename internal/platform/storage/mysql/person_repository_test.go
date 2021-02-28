package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	xmen "xmen-mutant/internal"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PersonRepository_Save_RepositoryError(t *testing.T) {
	personMutant, personDna := true, []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	person, err := xmen.NewPerson(personMutant, personDna)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO persons (mutant, dna) VALUES (?, ?)").
		WithArgs(personMutant, personDna).
		WillReturnError(errors.New("something-failed"))

	repo := NewPersonRepository(db, 1*time.Millisecond)

	_, err = repo.Save(context.Background(), person)

	assert.Error(t, err)
}

func Test_PersonRepository_Save_Succeed(t *testing.T) {
	/*
		personMutant, personDna := true, []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

		person, err := xmen.NewPerson(personMutant, personDna)
		require.NoError(t, err)

		db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)

		sqlMock.ExpectExec(
			"INSERT INTO persons (mutant, dna) VALUES (?, ?)").
			WithArgs(personMutant, personDna).
			WillReturnResult(sqlmock.NewResult(0, 1))

		repo := NewPersonRepository(db, 1*time.Second)

		_, err = repo.Save(context.Background(), person)

		assert.NoError(t, sqlMock.ExpectationsWereMet())
		assert.NoError(t, err)
	*/
}
