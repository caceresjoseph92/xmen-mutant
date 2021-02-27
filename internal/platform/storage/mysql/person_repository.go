package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	xmen "xmen-mutant/internal"

	"github.com/huandu/go-sqlbuilder"
)

// PersonRepository is a MySQL xmen.PersonRepository implementation.
type PersonRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewPersonRepository initializes a MySQL-based implementation of xmen.PersonRepository.
func NewPersonRepository(db *sql.DB, dbTimeout time.Duration) *PersonRepository {
	return &PersonRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the xmen.PersonRepository interface.
func (r *PersonRepository) Save(ctx context.Context, person xmen.Person) error {
	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	query, args := personSQLStruct.InsertInto(sqlPersonTable, sqlPerson{
		Mutant: person.Mutant().Bool(),
		Dna:    strings.Join(person.Dna().String(), ","),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist person on database: %v", err)
	}

	return nil
}
