package usecase

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"examenMutante/entity"
	mErrors "examenMutante/errors"
)

type (
	MutantsStore interface {
		InsertMutant(ctx context.Context, request entity.Response) error
		InsertHuman(ctx context.Context, request entity.Response) error
		GetStats(ctx context.Context) (*entity.Stats, error)
	}

	Mutants struct {
		Store MutantsStore
	}
)

var countSequence int

func NewMutants(mu MutantsStore) Mutants {
	return Mutants{
		Store: mu,
	}
}

func (mu Mutants) IsMutant(ctx context.Context, request entity.Request) (*entity.Response, error) {
	log := logrus.WithContext(ctx)
	dnaToInsert := entity.Response{
		Dna:  request.Dna,
		Name: request.Name,
	}
	if analyzeDNA(request.Dna) {
		err := mu.Store.InsertMutant(ctx, dnaToInsert)
		if err != nil {
			log.WithError(err).Errorf("Insert: error on db insert mutant")
			return nil, mErrors.Error{
				Cause:  err,
				Action: "cannot communicate with db on insert mutant",
				Status: 500,
				Code:   "69b21d90-f5bd-4ab2-b8db-c8935c9e324e",
			}
		}
		dnaToInsert.Type = "Mutant"
		return &dnaToInsert, nil
	}
	err := mu.Store.InsertHuman(ctx, dnaToInsert)
	if err != nil {
		log.WithError(err).Errorf("Insert: error on db insert humans")
		return nil, mErrors.Error{
			Cause:  err,
			Action: "cannot communicate with db on insert humans",
			Status: 500,
			Code:   "69b21d90-f5bd-4ab2-b8db-c8935c9e324e",
		}
	}
	dnaToInsert.Type = "Human"

	return nil, mErrors.Error{
		Cause:  fmt.Errorf("Not mutant"),
		Action: "DNA is not from a mutant",
		Status: 403,
		Code:   "9ea0438e-1b84-4e6e-91b5-2ed4dc311edf",
	}
}

func (mu Mutants) GetStats(ctx context.Context) (*entity.Stats, error) {
	log := logrus.WithContext(ctx)

	stats, err := mu.Store.GetStats(ctx)
	if err != nil {
		log.WithError(err).Errorf("GetStats: error on get stats")
		return nil, mErrors.Error{
			Cause:  err,
			Action: "cannot communicate with db to get stats",
			Status: 500,
			Code:   "8e11cdc3-ff1e-4b1b-8d12-0cd9651a62f8",
		}
	}

	return stats, err
}

func analyzeDNA(dna []string) bool {
	countSequence = 0
	adn := [6][6]string{}
	for a, values := range dna {
		newArray := strings.Split(values, "")
		for b, value := range newArray {
			adn[a][b] = value
		}
	}

	fmt.Println(" 	0 	1	2	3	4	5")
	fmt.Println(" -------------------------------------------------")
	fmt.Println("0	" + adn[0][0] + "	" + adn[0][1] + "	" + adn[0][2] + "	" + adn[0][3] + "	" + adn[0][4] + "	" + adn[0][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("1	" + adn[1][0] + "	" + adn[1][1] + "	" + adn[1][2] + "	" + adn[1][3] + "	" + adn[1][4] + "	" + adn[1][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("2	" + adn[2][0] + "	" + adn[2][1] + "	" + adn[2][2] + "	" + adn[2][3] + "	" + adn[2][4] + "	" + adn[2][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("3	" + adn[3][0] + "	" + adn[3][1] + "	" + adn[3][2] + "	" + adn[3][3] + "	" + adn[3][4] + "	" + adn[3][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("4	" + adn[4][0] + "	" + adn[4][1] + "	" + adn[4][2] + "	" + adn[4][3] + "	" + adn[4][4] + "	" + adn[4][5])
	fmt.Println(" -------------------------------------------------")
	fmt.Println("5	" + adn[5][0] + "	" + adn[5][1] + "	" + adn[5][2] + "	" + adn[5][3] + "	" + adn[5][4] + "	" + adn[5][5])
	fmt.Println(" -------------------------------------------------")

	horizontal(dna)
	vertical(dna)
	diagonal(dna)

	if countSequence >= 2 {
		return true
	}

	return false
}

func horizontal(dna []string) {
	for _, sequence := range dna {
		foundSequence(sequence)
	}
}

func vertical(dna []string) {
	for i := 0; i < len(dna); i++ {
		var sequence string
		for _, values := range dna {
			sequence += string(values[i])
		}
		foundSequence(sequence)
	}
}

func diagonal(dna []string) {
	var sequence string
	for i := 0; i < len(dna); i++ {
		for index, values := range dna {
			if i == index {
				sequence += string(values[i])
			}
		}
	}
	foundSequence(sequence)
}

func foundSequence(sequence string) {
	sequenceToFind := [4]string{"AAAA", "CCCC", "GGGG", "TTTT"}
	for _, n := range sequenceToFind {
		if sequence[0:4] == n {
			countSequence++
		}
	}
}
