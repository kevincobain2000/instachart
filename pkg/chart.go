package pkg

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/imroc/req/v3"
)

const (
	DEFAULT_PADDING_TOP        = 20
	DEFAULT_PADDING_RIGHT      = 20
	DEFAULT_PADDING_BOTTOM     = 20
	DEFAULT_PADDING_LEFT       = 20
	DEFAULT_SUBTITLE_FONT_SIZE = 10

	BAR_STYLE_VERTICAL   = "vertical"
	BAR_STYLE_HORIZONTAL = "horizontal"
	BAR_STYLE_STACKED    = "stacked"
)

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
	Fill          bool   `json:"fill" query:"fill" form:"fill" default:"false"`
}

func SetHeadersResponseImage(header http.Header) {
	header.Set("Cache-Control", "max-age=31536000")
	header.Set("Expires", "31536000")
	header.Set("Content-Type", "image/png")
	// security headers
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-XSS-Protection", "1; mode=block")
	// content policy
	header.Set("Content-Security-Policy", "default-src 'none'; img-src 'self'; style-src 'self'; font-src 'self'; connect-src 'self'; script-src 'self';")
}
func SetHeadersResponseHTML(header http.Header) {
	header.Set("Cache-Control", "max-age=31536000")
	header.Set("Expires", "31536000")
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

func IsURL(urlStr string) bool {
	parsedURL, err := url.ParseRequestURI(urlStr)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
func IsAllowedDomain(urlStr string, allowedDomains string) bool {
	if allowedDomains == "" {
		return false // default do not allow any urls
	}

	// Parse the URL to extract the domain
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false // If the URL is invalid, do not allow
	}
	domain := parsedURL.Hostname()

	// Split the allowedDomains into a slice
	domains := strings.Split(allowedDomains, ",")

	// Check if the domain is in the list of allowed domains
	for _, d := range domains {
		if domain == d {
			return true
		}
	}

	return false
}

func GetURL(urlStr string) (string, error) {
	resp, err := req.Get(urlStr)
	if err != nil {
		return "", err
	}
	return resp.ToString()
}

func SetDataIfRemoteURL(req *ChartRequest, allowedRemoteDomains string) error {
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
