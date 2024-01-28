package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mcuadros/go-defaults"
)

type TableChartHandler struct {
	chart *TableChart
}

func NewTableChartHandler() *TableChartHandler {
	return &TableChartHandler{
		chart: NewTableChart(),
	}
}

type TableChartData struct {
	Names  []string   `json:"names"`
	Values [][]string `json:"values"`
}

func (h *TableChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(ChartRequest)

	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	defaults.SetDefaults(req)
	msgs, err := ValidateRequest(req)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	err = SetDataIfRemoteURL(req)
	if err != nil {
		msgs := map[string]string{
			"data": err.Error(),
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	var data TableChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		msgs := map[string]string{
			"data": "Invalid JSON",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	if len(data.Names) == 0 {
		msgs := map[string]string{
			"data": "Counts are invalid",
		}
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}
	for _, row := range data.Values {
		if len(row) != len(data.Names) {
			msgs := map[string]string{
				"data": "Counts are invalid",
			}
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
		}
	}
	SetHeadersResponseImage(c.Response().Header(), req.Output)
	return h.chart.Get(data.Names, data.Values, req)
}
