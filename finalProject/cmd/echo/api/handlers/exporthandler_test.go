package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func TestExportList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/list/export", strings.NewReader(""))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	req.Header.Set("accept", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, ExportList(c)){
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t,  "Lists,tasks\n", rec.Body.String())
	}
}