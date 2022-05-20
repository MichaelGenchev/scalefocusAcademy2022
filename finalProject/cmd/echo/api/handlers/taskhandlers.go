package handlers

import (
	"final/cmd/echo/api/middlewares"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Task struct {
    Id int `json:"id"`
    Text string `json:"text"`
    ListId int `json:"listId"`
    Completed bool `json:"completed"`
}
var IDTask int = 0

func PatchTask(c echo.Context) error {
    
    updates := Task{}
    id := c.Param("taskId")
    err := c.Bind(&updates)
    if err != nil {
        log.Printf("Failed processing patchTask request: %s\n", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    db, _ := sqlx.Open("sqlite", middlewares.DBName)

    db.Exec("UPDATE tasks SET completed=? WHERE id=?", updates.Completed, id)
    
	task := Task{}

	taskRow, _ := db.Queryx("SELECT * FROM tasks WHERE id=?", id)

    for taskRow.Next() {
        err := taskRow.StructScan(&task)
        if err != nil {
            log.Fatalln(err)
        }
    }

    return c.JSON(http.StatusOK, task)
}

func PostTask(c echo.Context) error {

    task := Task{}

    parentListIdString := c.Param("id")
    parentListIdInt, _ := strconv.Atoi(parentListIdString) 

    err := c.Bind(&task)
    if err != nil {
        log.Printf("Failed processing postTask request: %s\n", err)
        return echo.NewHTTPError(http.StatusInternalServerError)
    }

    task.Id = IDTask
    IDTask = IDTask + 1
    task.ListId = parentListIdInt

    log.Printf("this is your task: %#v", task)

    addTaskToDatabase(task)

    return c.JSON(http.StatusOK, task)


}

func addTaskToDatabase(task Task){
    db, _ := sqlx.Open("sqlite", middlewares.DBName)

    db.Exec("INSERT INTO tasks (id, text, listid ,completed) VALUES ($1, $2, $3, $4)", task.Id, task.Text, task.ListId,task.Completed)
}

func DeleteTask(c echo.Context) error {
    taskId := c.Param("taskId")

    db, _ := sqlx.Open("sqlite", middlewares.DBName)

    db.Exec("DELETE FROM tasks WHERE id=?", taskId)

    return c.NoContent(http.StatusOK)
}

func GetTasks(c echo.Context) error {
    db, _ := sqlx.Open("sqlite", middlewares.DBName)

    tasks := []Task{}

    task := Task{}

    parentListId := c.Param("id")


    rows, err := db.Queryx("SELECT * FROM tasks WHERE listid=?", parentListId)
    if err != nil {
        log.Fatal("HERE", err)
    }
    for rows.Next() {
        rows.StructScan(&task)
        tasks = append(tasks, Task{task.Id, task.Text,task.ListId, task.Completed})
    }

    return c.JSON(http.StatusOK, tasks)
}

