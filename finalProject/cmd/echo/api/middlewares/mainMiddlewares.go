package middlewares

import (
	"log"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

const DBName = "database.db"
var IdUser int = 0
func SetMainMiddlewares(e *echo.Echo){
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        // This is a sample demonstration of how to attach middlewares in Echo
        return func(ctx echo.Context) error {
            log.Println("HOW MANY TIMES I WAS CALLED")
            return next(ctx)
        }
    })
	e.Use(middleware.CORS())

	e.Use(middleware.BasicAuth(func(username , password string , c echo.Context) (bool, error) {

		db, err := sqlx.Open("sqlite", DBName)
		if err != nil {
			log.Fatal(err)
		}
		// user := User{}
		
		rand.Seed(time.Now().Unix())
		currentId := rand.Int()
		hashedPassword, err := HashPassword(password)
		if err != nil {
			log.Fatal(err)
		}
		row := db.QueryRow("select username from users where username = ?", username)
		
		temp := ""
		row.Scan(&temp)
		if temp != "" {
			id := 0
			pass := ""
			rowId := db.QueryRow("select id from users where username = ?", username)
			rowPass := db.QueryRow("select password from users where username = ?", username)

			rowPass.Scan(&pass)
			rowId.Scan(&id)

			err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
			if err != nil {
				return false , nil
			}
			IdUser = id
			return true, nil
		}
		db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", currentId, username, hashedPassword)
		// c.Set("user_id", idUser)
		IdUser = currentId
		return true, nil
	}))
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
    return string(bytes), err
}
