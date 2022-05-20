package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)





var listJson = `{"id":0,"name":"test", "userid":214}`


func TestAddList(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/lists", strings.NewReader(listJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, AddList(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":0,\"name\":\"test\"}\n", rec.Body.String())
	}
}

func TestGetLists(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/lists", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	startDB()

	

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, GetLists(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestDeleteList(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/lists", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("authorization", "Basic cGV0ZXI6MTIz")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	startDB()

	

	res := rec.Result()
    defer res.Body.Close()

	if assert.NoError(t, DeleteList(c)){
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func startDB(){
    os.Remove("database.db")

    db, err := sqlx.Open("sqlite", "database.db")

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()


    _, err = db.Exec(`Create table lists (id INT PRIMARY KEY, name VARCHAR(32), userid INT)`)
    if err != nil {
        log.Print("From creating lists")
        log.Fatal(err)
    }
    _, err = db.Exec(`Create table tasks (id INT PRIMARY KEY, text VARCHAR(32),listid INT, completed bool)`)
    if err != nil {
        log.Print("From creating tasks")
        log.Fatal(err)
    }
	_, err = db.Exec(`Create table users (id INT PRIMARY KEY , username VARCHAR(32) , password VARCHAR(32))`)

    if err != nil {
        log.Print("From creating users")
        log.Fatal(err)
    }
}

