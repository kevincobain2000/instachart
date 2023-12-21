package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/line", func(c echo.Context) error {
		img, err := NewLineChartHandler().Get(c)
		if err != nil {
			return err
		}
		setHeaders(c)
		return c.Blob(http.StatusOK, "image/png", img)
	})
	e.GET("/bar", func(c echo.Context) error {
		img, err := NewBarChartHandler().Get(c)
		if err != nil {
			return err
		}
		setHeaders(c)
		return c.Blob(http.StatusOK, "image/png", img)
	})
}

func setHeaders(c echo.Context) {
	c.Response().Header().Set("Cache-Control", "max-age=31536000")
	c.Response().Header().Set("Expires", "31536000")
}
