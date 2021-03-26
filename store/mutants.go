package store

import (
	"context"
	"fmt"
	"github.com/lib/pq"

	"examenMutante/entity"
)

func (s Mutants) Insert(ctx context.Context, request entity.Response) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("Insert: cannot insert a mutant: %w", err)
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