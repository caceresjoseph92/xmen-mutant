package mysql

const (
	sqlPersonTable = "persons"
)

type sqlPerson struct {
	Mutant bool   `db:"mutant"`
	Dna    string `db:"dna"`
}
