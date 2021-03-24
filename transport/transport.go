package transport

import "github.com/labstack/echo"

func NewRouter(m Mutants, ) *echo.Echo {
	e := echo.New()

	mutants := e.Group("mutant")
	mutants.POST("", m.Create)
	return e
}