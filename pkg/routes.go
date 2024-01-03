package pkg

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	DOCS_URL     = "https://github.com/kevincobain2000/instachart"
	FAVICON_FILE = "favicon.ico"
	ROBOTS_FILE  = "robots.txt"
	ROBOTS_TXT   = `User-agent: *
Allow: *
Disallow: /line
Disallow: /bar
Disallow: /radar
Disallow: /funnel
Disallow: /donut
Disallow: /pie
Disallow: /table`
	DIST_DIR = "frontend/dist"
)

func SetupRoutes(e *echo.Echo, baseURL string, publicDir embed.FS, allowedRemoteDomains string) {

	e.GET(baseURL+"", func(c echo.Context) error {
		filename := fmt.Sprintf("%s/%s", DIST_DIR, "index.html")
		content, err := publicDir.ReadFile(filename)
		if err != nil {
			return c.String(http.StatusNotFound, "404")
		}
		SetHeadersResponseHTML(c.Response().Header())
		return c.Blob(http.StatusOK, "text/html", content)
	})

	// /robots.txt
	e.GET(baseURL+ROBOTS_FILE, func(c echo.Context) error {
		SetHeadersResponseTxt(c.Response().Header())
		return c.String(http.StatusOK, ROBOTS_TXT)
	})

	// /favicon.ico
	e.GET(baseURL+FAVICON_FILE, func(c echo.Context) error {
		filename := fmt.Sprintf("%s/%s", DIST_DIR, "favicon.ico")
		content, err := publicDir.ReadFile(filename)
		if err != nil {
			return c.String(http.StatusNotFound, "404")
		}
		SetHeadersResponseImage(c.Response().Header())
		return c.Blob(http.StatusOK, "image/x-icon", content)
	})

	// /line
	e.GET(baseURL+"line", func(c echo.Context) error {
		img, err := NewLineChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /bar
	e.GET(baseURL+"bar", func(c echo.Context) error {
		img, err := NewBarChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /radar
	e.GET(baseURL+"radar", func(c echo.Context) error {
		img, err := NewRadarChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /donut
	e.GET(baseURL+"donut", func(c echo.Context) error {
		img, err := NewDonutChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /pie
	e.GET(baseURL+"pie", func(c echo.Context) error {
		img, err := NewPieChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /funnel
	e.GET(baseURL+"funnel", func(c echo.Context) error {
		img, err := NewFunnelChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
	// /table
	e.GET(baseURL+"table", func(c echo.Context) error {
		img, err := NewTableChartHandler(allowedRemoteDomains).Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "", img)
	})
}
