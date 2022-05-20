package handlers

import (
	"encoding/csv"
	"final/cmd/echo/api/middlewares"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)




func ExportList(c echo.Context) error {

	resultRows := [][]string{
		{"Lists", "tasks", },
		
	}

	tasks := []Task{}
	list := List{}
	task := Task{}

	db, _ := sqlx.Open("sqlite", middlewares.DBName)

	rows, err := db.Queryx("SELECT * FROM lists WHERE userid=?", middlewares.IdUser)
    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        err := rows.StructScan(&list)
        if err != nil {
            log.Fatalln(err)
        }
		taskRows, err := db.Queryx("SELECT * FROM tasks WHERE listid=?", list.Id)
		if err != nil {
			log.Fatal("HERE", err)
		}
		for taskRows.Next() {
			err := taskRows.StructScan(&task)
			if err != nil {
				log.Fatalln(err)
			}
			tasks = append(tasks, Task{task.Id, task.Text,task.ListId, task.Completed})
		}
		resultRow := []string{list.Name}

		for _, task := range tasks {
			resultRow = append(resultRow, task.Text)
		}
		resultRows = append(resultRows, resultRow)
		tasks = []Task{}
    }


	csvfile, _ := os.Create("data.csv")

	cswriter := csv.NewWriter(csvfile)

	for _, row := range resultRows {
		_ = cswriter.Write(row)
	}
	cswriter.Flush()
	csvfile.Close()
	return c.Attachment("data.csv", "userdata")
}