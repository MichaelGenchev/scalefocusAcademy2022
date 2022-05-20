package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestGetTemperature(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/weather", strings.NewReader(""))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	req.Header.Set("accept", "application/json")
	req.Header.Set("lon", "23.319941")
	req.Header.Set("lat", "42.698334")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, GetTemperature(c)){
		assert.Equal(t, http.StatusOK, res.StatusCode)

	}
}