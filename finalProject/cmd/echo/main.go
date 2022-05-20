
package main

import (
	"final/cmd"
	"final/cmd/echo/api/middlewares"
	"final/cmd/echo/router"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "modernc.org/sqlite"


	"github.com/jmoiron/sqlx"
)


func main() {
    fmt.Println("Welcome to the server")
    startDB()

    e := router.New()

    log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(e)))
}



func startDB(){
    os.Remove(middlewares.DBName)

    db, err := sqlx.Open("sqlite", middlewares.DBName)

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