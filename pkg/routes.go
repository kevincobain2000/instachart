package pkg

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, baseURL string) {
	e.GET(baseURL+"health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET(baseURL+"robots.txt", func(c echo.Context) error {
		return c.String(http.StatusOK, "User-agent: *\nAllow: /")
	})
	e.GET(baseURL+"", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "https://github.com/kevincobain2000/instachart")
	})
	e.GET(baseURL+"line", func(c echo.Context) error {
		img, err := NewLineChartHandler().Get(c)
		if err != nil {
			return err
		}
		setHeaders(c)
		return c.Blob(http.StatusOK, "image/png", img)
	})
	e.GET(baseURL+"bar", func(c echo.Context) error {
		img, err := NewBarChartHandler().Get(c)
		if err != nil {
			return err
		}
		setHeaders(c)
		return c.Blob(http.StatusOK, "image/png", img)
	})
	e.GET(baseURL+"radar", func(c echo.Context) error {
		img, err := NewRadarChartHandler().Get(c)
		if err != nil {
			return err
		}
		setHeaders(c)
		return c.Blob(http.StatusOK, "image/png", img)
	})
	e.GET(baseURL+"donut", func(c echo.Context) error {
		img, err := NewDonutChartHandler().Get(c)
		if err != nil {
			return err
		}
		setHeaders(c)
		return c.Blob(http.StatusOK, "image/png", img)
	})
	e.GET(baseURL+"pie", func(c echo.Context) error {
		img, err := NewPieChartHandler().Get(c)
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
