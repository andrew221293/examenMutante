package transport

import (
	"context"
	"examenMutante/entity"
	"github.com/labstack/echo"
	"src.srconnect.io/pkg/ctxerr"
)

type MutantsUsacase interface {
	ValidateMutant(ctx context.Context, request entity.MutantsRequest) (*entity.MutantsRequest, error)
}

type Mutants struct {
	Usecase MutantsUsacase
}

func NewMutant(m MutantsUsacase) Mutants {
	return Mutants{Usecase: m}
}

func(mt Mutants) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var mutantRequest entity.MutantsRequest
	if err := c.Bind(&mutantRequest); err != nil {
		ctx := ctxerr.SetHTTPStatusCode(ctx, 400)
		ctx = ctxerr.SetAction(ctx, "invalidBody")

		return ctxerr.Wrap(ctx, err, "64bd031a-8e17-4896-a79d-32e01be5643e", "Request body is invalid")
	}

	mutansCreated, err := mt.Usecase.ValidateMutant(ctx, mutantRequest)
	if err != nil {
		return err
	}

	return c.JSON(201, mutansCreated)
}