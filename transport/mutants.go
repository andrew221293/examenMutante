package transport

import (
	"context"

	"examenMutante/entity"
	mErrors "examenMutante/errors"

	"github.com/labstack/echo"
)

type (
	MutantsUseCase interface {
		IsMutant(ctx context.Context, request entity.Request) (*entity.Response, error)
	}
	Mutants struct {
		UseCase MutantsUseCase
	}
)

func NewMutant(m MutantsUseCase) Mutants {
	return Mutants{UseCase: m}
}

func (mt Mutants) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var mutantRequest entity.Request
	err := c.Bind(&mutantRequest)
	if err != nil {
		return mErrors.Error{
			Cause:  err,
			Action: "couldn't parse body ",
			Status: 404,
			Code:   "64d669af-11ac-4716-b060-720a23f461ed",
		}
	}

	isMutant, err := mt.UseCase.IsMutant(ctx, mutantRequest)
	if err != nil {
		return err
	}

	return c.JSON(200, isMutant)
}
