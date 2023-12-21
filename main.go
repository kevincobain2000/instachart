package main

import (
	"flag"
	"net/http"

	"github.com/kevincobain2000/go-chart-to-picture/pkg"
	"github.com/labstack/echo/v4"
)

var (
	port string
)

func main() {
	cliArgs()
	e := pkg.NewEcho()

	e.GET("/line", func(c echo.Context) error {
		img, err := pkg.NewLineChartHandler().GetLineRequestChart(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "image/png", img)
	})

	pkg.GracefulServerWithPid(e, port)
}

func cliArgs() {
	flag.StringVar(&port, "port", "3000", "port to serve")
	flag.Parse()
}
