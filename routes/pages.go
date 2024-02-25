package routes

import (
	"github.com/JakubMrowicki/echo-go-web-server/components"
	"github.com/JakubMrowicki/echo-go-web-server/services"
	"github.com/labstack/echo/v4"
)

func Pages(e *echo.Echo) {
	e.File("/", "assets/public/index.html")
	e.GET("/weather", func(c echo.Context) error {
		temp := services.RandomWeather() + "°C"
		component := components.Boilerplate(temp)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
}
