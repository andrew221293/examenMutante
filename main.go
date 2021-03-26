package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	mErrors "examenMutante/errors"
	"examenMutante/transport"
	"examenMutante/usecase"

	"github.com/labstack/echo/middleware"
)

func main() {
	mutantsUseCase := usecase.NewMutants()

	mutantsTransport := transport.NewMutant(mutantsUseCase)

	e := transport.NewRouter(mutantsTransport)
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = mErrors.EchoErrorHandler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Info("shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
