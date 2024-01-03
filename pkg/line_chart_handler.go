package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LineChartHandler struct {
	chart                *LineChart
	allowedRemoteDomains string
}

func NewLineChartHandler(allowedRemoteURLS string) *LineChartHandler {
	return &LineChartHandler{
		chart:                NewLineChart(),
		allowedRemoteDomains: allowedRemoteURLS,
	}
}

type LineChartData struct {
	XData []string    `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *LineChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	err := SetDataIfRemoteURL(req, h.allowedRemoteDomains)
	if err != nil {
		msgs := map[string]string{
			"data": err.Error(),
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	var data LineChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	if len(data.XData) == 0 || len(data.YData) == 0 {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	for _, y := range data.YData {
		if len(y) != len(data.XData) {
			msgs := map[string]string{
				"data": "Counts are invalid",
			}
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
		}
	}

	SetHeadersResponseImage(c.Response().Header())
	return h.chart.Get(data.XData, data.YData, data.Names, req)
}
