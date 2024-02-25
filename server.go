package main

import (
	"github.com/JakubMrowicki/echo-go-web-server/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "assets/public",
		Browse: true,
	}))

	// Page Routes
	routes.Pages(e)
	routes.Api(e)

	e.Logger.Fatal(e.Start(":3000"))
}
