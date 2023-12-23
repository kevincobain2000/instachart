package pkg

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fvbock/endless"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = HTTPErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "REQ [${time_custom}] ${method} ${uri} (${latency_human}) ${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	SetupRoutes(e)
	return e
}

// GracefulServerWithPid reloads server with pid
// kill -HUP when binary is changed
// kill -9 when want to kill the process and make the application dead and want to restart
// kill -9 is NOT FOR FAINT HEARTED and must not be done on prod unless SOUT
func GracefulServerWithPid(e *echo.Echo, port string) {
	log := Logger()
	server := endless.NewServer("localhost:"+port, e)
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
	if err = c.JSON(code, &HTTPErrorResponse{Error: message}); err != nil {
		Logger().Error(err)
	}
}
