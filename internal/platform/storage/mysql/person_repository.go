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

var (
	count_mutant_dna int
	count_human_dna  int
	ratio            float64
)

type row struct {
	value int
}

// NewPersonRepository initializes a MySQL-based implementation of xmen.PersonRepository.
func NewPersonRepository(db *sql.DB, dbTimeout time.Duration) *PersonRepository {
	return &PersonRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the xmen.PersonRepository interface.
func (r *PersonRepository) Save(ctx context.Context, person xmen.Person) (persons map[string]interface{}, errs error) {
	personSQLStruct := sqlbuilder.NewStruct(new(sqlPerson))
	query, args := personSQLStruct.InsertInto(sqlPersonTable, sqlPerson{
		Mutant: person.Mutant().Bool(),
		Dna:    strings.Join(person.Dna().String(), ","),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error trying to persist person on database: %v", err)
	}

	return nil, nil
}

// Consult implements the xmen.PersonRepository interface.
func (r *PersonRepository) Consult(ctx context.Context, args map[string]interface{}) (stats map[string]interface{}, err error) {
	got := []row{}
	rowsMutant, errs := r.db.QueryContext(ctx, "SELECT COUNT(*) AS mutant FROM persons WHERE mutant=?", true)
	if errs != nil {
		return nil, fmt.Errorf("error trying to struct query on database: %v", errs)
	}
	rowsHuman, errs := r.db.QueryContext(ctx, "SELECT COUNT(*) AS mutant FROM persons WHERE mutant=?", false)
	if errs != nil {
		return nil, fmt.Errorf("error trying to struct query on database: %v", errs)
	}
	for rowsMutant.Next() {
		if err := rowsMutant.Err(); err != nil {
			return nil, fmt.Errorf("error trying to struct query on database: %v", err)
		}
		var ra row
		rowsMutant.Scan(&ra.value)
		got = append(got, ra)
	}
	for rowsHuman.Next() {
		if err := rowsHuman.Err(); err != nil {
			return nil, fmt.Errorf("error trying to struct query on database: %v", err)
		}
		var ra row
		rowsHuman.Scan(&ra.value)
		got = append(got, ra)
	}

	for i, row := range got {
		if i == 0 {
			count_mutant_dna = row.value
		} else {
			count_human_dna = row.value
		}
	}
	ratio = float64(count_mutant_dna) / float64(count_human_dna)

	stats = map[string]interface{}{
		"count_mutant_dna": count_mutant_dna,
		"count_human_dna":  count_human_dna,
		"ratio":            ratio,
	}

	return stats, nil
}
