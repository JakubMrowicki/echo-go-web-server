package main

import (
	"fmt"
	"net/http"

	"github.com/JakubMrowicki/echo-go-web-server/components"
	"github.com/JakubMrowicki/echo-go-web-server/services"
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

	// Routes
	e.File("/", "assets/public/index.html")
	e.GET("/weather", func(c echo.Context) error {
		temp := services.RandomWeather() + "°C"
		component := components.Boilerplate(temp)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.GET("/get-temperature", func(c echo.Context) error {
		temp := services.RandomWeather()
		return c.String(http.StatusOK, fmt.Sprintf("%v °C", temp))
	})
	e.Any("/*", func(c echo.Context) error {
		c.Error(echo.ErrNotFound)
		return nil
	})

	e.Logger.Fatal(e.Start(":3000"))
}
