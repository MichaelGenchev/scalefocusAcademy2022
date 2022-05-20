package api

import (
	"net/http"
	"final/cmd/echo/api/handlers"


	"github.com/labstack/echo/v4"
)


func MainGroup(e *echo.Echo) {
	e.GET("/api/", func(ctx echo.Context) error {
        return ctx.JSON(http.StatusOK, "Hello, world!")
    })

    // CSV FILE
	e.GET("/api/list/export", handlers.ExportList)

    // LISTS
    e.DELETE("/api/lists/:id", handlers.DeleteList)
    e.GET("/api/lists", handlers.GetLists)
    e.POST("/api/lists", handlers.AddList)


    //TASKS    

    e.POST("/api/lists/:id/tasks", handlers.PostTask)
    e.GET("/api/lists/:id/tasks", handlers.GetTasks)
    e.DELETE("/api/tasks/:taskId", handlers.DeleteTask)
    e.PATCH("/api/tasks/:taskId", handlers.PatchTask)

    // WEATHER
    e.GET("/api/weather", handlers.GetTemperature)
}