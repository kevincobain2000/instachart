package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	charts "github.com/vicanso/go-charts/v2"
)

type BarChartHandler struct {
	chart *BarChart
}

func NewBarChartHandler() *BarChartHandler {
	return &BarChartHandler{
		chart: NewBarChart(),
	}
}

type BarChartRequest struct {
	ChartData     string  `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle    string  `json:"title" query:"title" form:"title"`
	ChartSubtitle string  `json:"subtitle" query:"subtitle" form:"subtitle"`
	BaseValue     float64 `json:"base_value" query:"base_value" form:"base_value"`
	Height        int     `json:"height" query:"height" form:"height"`
	Width         int     `json:"width" query:"width" form:"width"`
	Horizontal    bool    `json:"horizontal" query:"horizontal" form:"horizontal"`
}

type BarChartData struct {
	XData []string    `json:"x"`
	YData [][]float64 `json:"y"`
	Names []string    `json:"names"`
}

func (h *BarChartHandler) Get(c echo.Context) ([]byte, error) {
	req := new(BarChartRequest)
	if err := BindRequest(c, req); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var data BarChartData
	if err := json.Unmarshal([]byte(req.ChartData), &data); err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if len(data.XData) == 0 || len(data.XData) != len(data.YData[0]) {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, "data provided is invalid")
	}
	if len(data.YData) == 1 && !req.Horizontal {
		// append [0,0,...0] to data.YData
		data.YData = append(data.YData, make([]float64, len(data.YData[0])))
	}

	if !req.Horizontal {
		p, err := charts.BarRender(
			data.YData,
			charts.TitleOptionFunc(charts.TitleOption{
				Text:            req.ChartTitle,
				Subtext:         req.ChartSubtitle,
				SubtextFontSize: 9,
				Left:            charts.PositionCenter,
			}),
			charts.HeightOptionFunc(req.Height),
			charts.WidthOptionFunc(req.Width),
			charts.XAxisDataOptionFunc(data.XData),
			charts.LegendLabelsOptionFunc(data.Names, charts.PositionRight),
			charts.MarkLineOptionFunc(0, charts.SeriesMarkDataTypeAverage),
			charts.MarkPointOptionFunc(0, charts.SeriesMarkDataTypeMax,
				charts.SeriesMarkDataTypeMin),
			func(opt *charts.ChartOption) {
				opt.SeriesList[1].MarkPoint = charts.NewMarkPoint(
					charts.SeriesMarkDataTypeMax,
					charts.SeriesMarkDataTypeMin,
				)
				opt.SeriesList[1].MarkLine = charts.NewMarkLine(
					charts.SeriesMarkDataTypeAverage,
				)
			},
		)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		buf, err := p.Bytes()
		return buf, err
	}

	p, err := charts.HorizontalBarRender(
		data.YData,
		charts.TitleOptionFunc(charts.TitleOption{
			Text:            req.ChartTitle,
			Subtext:         req.ChartSubtitle,
			SubtextFontSize: 9,
			Left:            charts.PositionCenter,
		}),
		charts.HeightOptionFunc(req.Height),
		charts.WidthOptionFunc(req.Width),
		charts.YAxisDataOptionFunc(data.XData),
	)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	buf, err := p.Bytes()
	return buf, err
}
