package store

import (
	"database/sql"
	"fmt"
	"time"
)

type (
	Mutants struct {
		db *sql.DB
	}
)

func NewStore(db *sql.DB, conMaxLifeTime int) (Mutants, error) {
	if err := db.Ping(); err != nil {
		return Mutants{}, fmt.Errorf("could not ping postgres database: %v", err)
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(conMaxLifeTime))
	return Mutants{
		db: db,
	}, nil
}
