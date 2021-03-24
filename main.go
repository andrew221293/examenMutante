package main

import (
	"examenMutante/transport"
	"examenMutante/usecase"
	"github.com/labstack/echo/middleware"
	_ "src.srconnect.io/pkg/logger"
	_ "src.srconnect.io/pkg/logger/echologger"
)

func main() {

	// Routes
	mutantsUsecase := usecase.NewMutants()
	mutantsTransport := transport.NewMutant(mutantsUsecase)

	e := transport.NewRouter(mutantsTransport)
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

