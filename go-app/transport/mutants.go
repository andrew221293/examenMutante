package transport

import (
	"context"

	"github.com/labstack/echo"

	"Mutants/go-app/entity"
	mErrors "Mutants/go-app/errors"
)

type (
	MutantsUseCase interface {
		IsMutant(ctx context.Context, request entity.Request) (*entity.Response, error)
		GetStats(ctx context.Context) (*entity.Stats, error)
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

func (mt Mutants) GetStats(c echo.Context) error {
	ctx := c.Request().Context()

	stats, err := mt.UseCase.GetStats(ctx)
	if err != nil {
		return err
	}

	return c.JSON(200, stats)
}
