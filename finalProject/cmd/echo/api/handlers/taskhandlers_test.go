package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var taskJson = `{"id":0,"text":"test", "listid":0}`
var updateJson = `{"completed":true}`

func TestPostTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(taskJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/lists/:id/tasks")
	c.SetParamNames("id")
	c.SetParamValues("0")

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, PostTask(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":0,\"text\":\"test\",\"listId\":0,\"completed\":false}\n", rec.Body.String())
	}
}

func TestDeleteTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(taskJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/tasks/:taskId")
	c.SetParamNames("taskId")
	c.SetParamValues("0")

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, DeleteTask(c)){
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestPatchTask(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/lists/:id/tasks")
	c.SetParamNames("id")
	c.SetParamValues("0")

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, PatchTask(c)){
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t,  "{\"id\":0,\"text\":\"\",\"listId\":0,\"completed\":false}\n", rec.Body.String())
	}
}


func TestGetTasks(t *testing.T){
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(updateJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/tasks/:taskId")
	c.SetParamNames("taskId")
	c.SetParamValues("0")

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, GetTasks(c)){
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t,  "[]\n", rec.Body.String())
	}

}
