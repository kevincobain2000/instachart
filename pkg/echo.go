package pkg

import (
	"embed"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho(baseURL string, publicDir embed.FS) *echo.Echo {
	if os.Getenv("PPROF_HOST") != "" && os.Getenv("PPROF_PORT") != "" {
		Logger().Info("pprof enabled and listening on: ", os.Getenv("PPROF_HOST")+":"+os.Getenv("PPROF_PORT"))
		addr := os.Getenv("PPROF_HOST") + ":" + os.Getenv("PPROF_PORT")
		viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeWesteros), viewer.WithAddr(addr))
		mgr := statsview.New()
		_ = mgr
		go mgr.Start()
		// mgr.Stop()
	}
	e := echo.New()
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = HTTPErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Gzip())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "REQUEST[${time_custom}] ${method} ${uri} (${latency_human}) ${status} ${remote_ip}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	SetupRoutes(e, baseURL, publicDir)
	return e
}

// GracefulServerWithPid reloads server with pid
// kill -HUP when binary is changed
// kill -9 when want to kill the process and make the application dead and want to restart
// kill -9 is NOT FOR FAINT HEARTED and must not be done on prod unless SOUT
func GracefulServerWithPid(e *echo.Echo, host string, port string) {
	log := Logger()
	server := endless.NewServer(host+":"+port, e)
	server.BeforeBegin = func(add string) {
		pidFile := filepath.Join(port + ".pid")
		_ = os.Remove(pidFile)
		err := os.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0600)
		if err != nil {
			log.Error("write pid file error: ", err)
		}
		log.Info("started server on localhost:", port)
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error("graceful error: ", err)
	}
}

// HTTPErrorResponse is the response for HTTP errors
type HTTPErrorResponse struct {
	Error interface{} `json:"error"`
}

// HTTPErrorHandler handles HTTP errors for entire application
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		message = err.Error()
	}

	Logger().Error(message)
	if code == http.StatusInternalServerError {
		message = "Internal Server Error"
	}
	if err = c.JSON(code, &HTTPErrorResponse{Error: message}); err != nil {
		Logger().Error(err)
	}
}
