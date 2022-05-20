package handlers

import (
	"log"
	"net/http"
	"final/cmd/echo/api/middlewares"
    _ "modernc.org/sqlite"


	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type List struct {
    Id int  `json:"id"`
    Name string `json:"name"`
	UserId int `json:"userid"`
}
type ResultList struct {
    Id int  `json:"id"`
    Name string `json:"name"`
}
var IDList int = 0


func GetLists(c echo.Context) error {

    db, err := sqlx.Open("sqlite", middlewares.DBName)
    if err != nil {
        log.Fatal(err)
    }
    lists := []List{}

    list := List{}
	

    rows, err := db.Queryx("SELECT * FROM lists WHERE userid=?", middlewares.IdUser)
    if err != nil {
        log.Fatal(err)
    }
    for rows.Next() {
        rows.StructScan(&list)

        lists = append(lists, List{list.Id, list.Name,list.UserId})
    }

	result := []ResultList{}
	for _, list := range lists {
		resultSingleList := ResultList{}
		resultSingleList.Id = list.Id
		resultSingleList.Name = list.Name

		result = append(result, resultSingleList)
	}

    return c.JSON(http.StatusOK, result)

}

func DeleteList(ctx echo.Context) error {
    id := ctx.Param("id")
    db, err := sqlx.Open("sqlite", middlewares.DBName)
    if err != nil {
        log.Fatal("failed to connect database")
    }
    db.Exec("DELETE FROM lists WHERE id=?", id)
    return ctx.NoContent(http.StatusOK)
}

func AddList(c echo.Context) error {
    list := List{}
    err := c.Bind(&list)
    if err != nil {
        log.Printf("Failed processing addList request: %s\n", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    list.Id = IDList
	list.UserId = middlewares.IdUser
    IDList = IDList + 1

    log.Printf("this is your list: %#v", list)

    addListToDatabase(list)

	resultList := ResultList{}
	resultList.Id = list.Id
	resultList.Name = list.Name


    return c.JSON(http.StatusOK, resultList)
}

func addListToDatabase(list List) {
    db, _ := sqlx.Open("sqlite", middlewares.DBName)

    db.Exec("INSERT INTO lists (id, name, userid) VALUES ($1, $2, $3)", list.Id, list.Name, list.UserId)
}
