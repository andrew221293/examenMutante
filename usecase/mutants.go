package usecase

import (
	"context"
	"fmt"
	"strings"

	"examenMutante/entity"
)

type (
	MutantsStore interface {
		Insert(ctx context.Context, request entity.Request) error
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
	if analyzeDNA(request.Dna) {
		return &entity.Response{
			request.Dna,
			request.Name,
			"Mutant",
		}, nil
	}
	return &entity.Response{
		request.Dna,
		request.Name,
		"Human",
	}, nil
}

func analyzeDNA(dna [6]string) bool {
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

func horizontal(dna [6]string) {
	for _, sequence := range dna {
		foundSequence(sequence)
	}
}

func vertical(dna [6]string) {
	for i := 0; i < len(dna); i++ {
		var sequence string
		for _, values := range dna {
			sequence += string(values[i])
		}
		foundSequence(sequence)
	}
}

func diagonal(dna [6]string) {
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
			countSequence ++
		}
	}
}
