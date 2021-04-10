package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"

	"Mutants/go-app/entity"
)

func (s Mutants) InsertMutant(ctx context.Context, request entity.Response) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("Insert: cannot insert a mutant: %w", err)
	}

	if _, err = tx.ExecContext(ctx,createTableMutants); err != nil {
		log.Fatalf("DB.ExecContext: unable to create table: %s", err)
	}

	_, err = tx.ExecContext(ctx, mutantStatement,
		pq.Array(request.Dna),
		request.Name,
	)
	if err != nil {
		secondErr := tx.Rollback()
		if secondErr != nil {
			err = fmt.Errorf("%v : %w", err, secondErr)
		}
		return err
	}

	return tx.Commit()
}

func (s Mutants) InsertHuman(ctx context.Context, request entity.Response) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("Insert: cannot insert a human: %w", err)
	}

	if _, err = tx.ExecContext(ctx,createTableHumans); err != nil {
		log.Fatalf("DB.ExecContext: unable to create table: %s", err)
	}

	_, err = tx.ExecContext(ctx, humanStatement,
		pq.Array(request.Dna),
		request.Name,
	)
	if err != nil {
		secondErr := tx.Rollback()
		if secondErr != nil {
			err = fmt.Errorf("%v : %w", err, secondErr)
		}
		return err
	}

	return tx.Commit()
}

func (s Mutants) GetStats(ctx context.Context) (*entity.Stats, error) {
	statsMutant := s.db.QueryRowContext(ctx, mutantsTotalStatement)
	statsHuman := s.db.QueryRowContext(ctx, humansTotalStatement)

	stats := &entity.Stats{}
	err := statsMutant.Scan(
		&stats.CountMutantDna,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("could't count mutants: %w", "mutants could not be counted")
		}
		return nil, err
	}

	errSecond := statsHuman.Scan(
		&stats.CountHumanDna,
	)
	if errSecond != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("could't count humans: %w", "humans could not be counted")
		}
		return nil, err
	}

	stats.Ratio = float32(stats.CountMutantDna) / float32(stats.CountHumanDna)

	return stats, nil

}
