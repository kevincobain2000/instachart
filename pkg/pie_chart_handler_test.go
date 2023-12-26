package pkg

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetPieChart(t *testing.T) {
	e := echo.New()

	e.GET("/pie", func(c echo.Context) error {
		img, err := NewPieChartHandler().Get(c)
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
		url := server.URL + "/pie?data=" + url.QueryEscape(tc.QueryParams)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedStatus, resp.StatusCode)
	}
}
