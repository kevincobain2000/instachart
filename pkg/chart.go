package pkg

import (
	"errors"
	"net/http"
	"os"

	charts "github.com/vicanso/go-charts/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

const (
	DEFAULT_PADDING_TOP        = 20
	DEFAULT_PADDING_RIGHT      = 20
	DEFAULT_PADDING_BOTTOM     = 20
	DEFAULT_PADDING_LEFT       = 20
	DEFAULT_TITLE_FONT_SIZE    = 12
	DEFAULT_SUBTITLE_FONT_SIZE = 10
	MINI_CHART_WIDTH           = 300
	MINI_CHART_HEIGHT          = 300

	BAR_STYLE_VERTICAL   = "vertical"
	BAR_STYLE_HORIZONTAL = "horizontal"
	BAR_STYLE_STACKED    = "stacked"
)

// slate
var DEFAULT_SUBTITLE_COLOR = drawing.Color{
	R: 112,
	G: 128,
	B: 144,
	A: 255,
}

var DEFAULT_BACKGROUND_COLOR = drawing.Color{
	R: 0,
	G: 0,
	B: 0,
	A: 0,
}

type Chart struct {
}

func NewChart() *Chart {
	return &Chart{}
}

type ChartRequest struct {
	ChartData     string `json:"data" query:"data" form:"data" validate:"required" message:"data is required"`
	ChartTitle    string `json:"title" query:"title" form:"title"`
	ChartSubtitle string `json:"subtitle" query:"subtitle" form:"subtitle"`
	Metric        string `json:"metric" query:"metric" form:"metric"`
	ZMetric       string `json:"zmetric" query:"zmetric" form:"zmetric"`
	Theme         string `json:"theme" query:"theme" form:"theme" default:"light"`
	Width         int    `json:"width" query:"width" form:"width" default:"1024"`
	Height        int    `json:"height" query:"height" form:"height" default:"768"`
	Style         string `json:"style" query:"style" form:"style" default:"vertical"`
	Line          string `json:"line" query:"line" default:"nofill" validate:"oneof=nofill fill" message:"line must be fill"`
	Output        string `json:"output" query:"output" form:"output" default:"png"`
}

func SetHeadersResponseImage(header http.Header, output string) {
	header.Set("Cache-Control", "max-age=31536000")
	header.Set("Expires", "31536000")
	if output == "svg" {
		header.Set("Content-Type", "image/svg+xml; charset=utf-8")
	} else {
		header.Set("Content-Type", "image/png")
	}
	// // security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// Strict-Transport-Security
	header.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
}
func SetHeadersResponseHTML(header http.Header) {
	header.Set("Cache-Control", "max-age=86400")
	header.Set("Expires", "86400")
	header.Set("Content-Type", "text/html; charset=utf-8")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
}

func SetHeadersResponseTxt(header http.Header) {
	header.Set("Cache-Control", "max-age=86400")
	header.Set("Expires", "86400")
	header.Set("Content-Type", "text/plain; charset=utf-8")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
}

func SetDataIfRemoteURL(req *ChartRequest) error {
	allowedRemoteDomains := os.Getenv("ALLOWED_REMOTE_DOMAINS")
	if allowedRemoteDomains == "" {
		return nil
	}
	if IsURL(req.ChartData) {
		if !IsAllowedDomain(req.ChartData, allowedRemoteDomains) {
			return errors.New("URL is not allowed")
		}
		data, err := GetURL(req.ChartData)
		if err != nil {
			return err
		}
		req.ChartData = string(data)
	}
	return nil
}

func IsMiniChart(req *ChartRequest) bool {
	return req.Width <= MINI_CHART_WIDTH && req.Height <= MINI_CHART_HEIGHT
}

func GetPaddings(req *ChartRequest) charts.Box {
	paddings := charts.Box{
		Top:    10,
		Bottom: 10,
		Left:   10,
		Right:  10,
	}
	if IsMiniChart(req) {
		paddings = charts.Box{
			Top:    10,
			Bottom: -20,
			Left:   -10,
			Right:  10,
		}
	}
	return paddings
}

func GetTitleSizes(req *ChartRequest) charts.TitleOption {
	titleSizes := charts.TitleOption{
		Text:             req.ChartTitle,
		Subtext:          req.ChartSubtitle,
		FontSize:         DEFAULT_TITLE_FONT_SIZE,
		SubtextFontSize:  DEFAULT_SUBTITLE_FONT_SIZE,
		Left:             charts.PositionCenter,
		SubtextFontColor: DEFAULT_SUBTITLE_COLOR,
	}
	if IsMiniChart(req) {
		titleSizes = charts.TitleOption{
			Text:             Truncate(req.ChartTitle, 17),
			Subtext:          Truncate(req.ChartSubtitle, 17),
			FontSize:         DEFAULT_TITLE_FONT_SIZE,
			SubtextFontSize:  DEFAULT_SUBTITLE_FONT_SIZE,
			Left:             charts.PositionCenter,
			SubtextFontColor: DEFAULT_SUBTITLE_COLOR,
		}
	}
	return titleSizes
}
