package store

import (
	"database/sql"
	"fmt"
)

type (
	Mutants struct {
		db *sql.DB
	}
)

func NewStore(db *sql.DB) (Mutants, error) {
	if err := db.Ping(); err != nil {
		return Mutants{}, fmt.Errorf("could not ping postgres database: %v", err)
	}
	return Mutants{
		db: db,
	}, nil
}
