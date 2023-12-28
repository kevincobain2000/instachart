package pkg

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed favicon.ico
var faviconEmbed embed.FS

const (
	DOCS_URL     = "https://github.com/kevincobain2000/instachart"
	FAVICON_FILE = "favicon.ico"
	ROBOTS_FILE  = "robots.txt"
	ROBOTS_TXT   = "User-agent: *\nAllow: /"
)

func SetupRoutes(e *echo.Echo, baseURL string) {
	fav, _ := faviconEmbed.ReadFile(FAVICON_FILE)

	// /
	e.GET(baseURL+"", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, DOCS_URL)
	})

	// /health
	e.GET(baseURL+"health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	// /robots.txt
	e.GET(baseURL+ROBOTS_FILE, func(c echo.Context) error {
		return c.String(http.StatusOK, ROBOTS_TXT)
	})

	// /favicon.ico
	e.GET(baseURL+FAVICON_FILE, func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/x-icon", fav)
	})

	// /line
	e.GET(baseURL+"line", func(c echo.Context) error {
		img, err := NewLineChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /bar
	e.GET(baseURL+"bar", func(c echo.Context) error {
		img, err := NewBarChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /radar
	e.GET(baseURL+"radar", func(c echo.Context) error {
		img, err := NewRadarChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /donut
	e.GET(baseURL+"donut", func(c echo.Context) error {
		img, err := NewDonutChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /pie
	e.GET(baseURL+"pie", func(c echo.Context) error {
		img, err := NewPieChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /funnel
	e.GET(baseURL+"funnel", func(c echo.Context) error {
		img, err := NewFunnelChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
}
