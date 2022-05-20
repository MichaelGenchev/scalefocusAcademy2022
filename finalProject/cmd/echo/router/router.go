package router

import (
	"final/cmd/echo/api/middlewares"
	"final/cmd/echo/api"


	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	router := echo.New()

	// SET MIDDLEWARES
	middlewares.SetMainMiddlewares(router)

	// SET MAIN ROUTES
	api.MainGroup(router)

	return router
}