package pkg

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetDonutChart(t *testing.T) {
	e := echo.New()

	e.GET("/donut", func(c echo.Context) error {
		img, err := NewDonutChartHandler().Get(c)
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "image/png", img)
	})

	// Start a test HTTP server
	server := httptest.NewServer(e)
	defer server.Close()

	type TestCase struct {
		QueryParams    string
		ExpectedStatus int
	}
	testCases := []TestCase{
		{
			QueryParams:    `{"names": ["Monday","Tuesday","Wednesday"], "values": [1,2,3]}`,
			ExpectedStatus: http.StatusOK,
		},
		{
			QueryParams:    `{"names": ["Monday,"Tuesday","Wednesday"], "values": [1,2,3]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			QueryParams:    `{"names": [["Monday","Tuesday","Wednesday"]], "values": [1,2,3]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		url := server.URL + "/donut?data=" + url.QueryEscape(tc.QueryParams)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedStatus, resp.StatusCode)
		if resp.StatusCode == http.StatusOK {
			assert.Equal(t, "image/png", resp.Header.Get("Content-Type"))
		} else {
			assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))
		}
	}
}
