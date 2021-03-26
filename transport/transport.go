package transport

import "github.com/labstack/echo"

func NewRouter(m Mutants) *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"Hello": "Test Mutant",
		})
	})

	mutants := e.Group("mutant")
	mutants.POST("", m.Create)

	return e
}
