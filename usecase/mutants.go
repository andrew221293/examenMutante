package usecase

import (
	"context"
	"examenMutante/entity"
	"fmt"
)

type Mutant struct {

}
func NewMutants() Mutant {
	return Mutant{

	}
}

func (mu Mutant)ValidateMutant(ctx context.Context, request entity.MutantsRequest) (*entity.MutantsRequest, error) {
	fmt.Println(ctx)
	fmt.Println(request)
return &entity.MutantsRequest{}, nil
}

