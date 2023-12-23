package pkg

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetLineChart(t *testing.T) {
	e := echo.New()

	e.GET("/line", func(c echo.Context) error {
		img, err := NewLineChartHandler().Get(c)
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
			QueryParams:    `{"x": [["2022-12-23","2022-12-24","2023-12-25"], ["2022-12-23","2022-12-28","2023-12-30"]], "y": [[1,2,3], [2,5,3]]}`,
			ExpectedStatus: http.StatusOK,
		},
		{
			QueryParams:    `{"x": [["abc","2022-12-24","2023-12-25"], ["2022-12-23","2022-12-28","2023-12-30"]], "y": [[1,2,3], [2,5,3]]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
		{
			QueryParams:    `{"x": [[2022-12-23,"2022-12-24","2023-12-25"], ["2022-12-23","2022-12-28","2023-12-30"]], "y": [[1,2,3], [2,5,3]]}`,
			ExpectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		url := server.URL + "/line?data=" + url.QueryEscape(tc.QueryParams)
		resp, err := http.Get(url)
		assert.NoError(t, err)
		assert.Equal(t, tc.ExpectedStatus, resp.StatusCode)

		// Additional checks for response body, etc.
	}
}
