package routes

import (
	"fmt"
	"net/http"

	"github.com/JakubMrowicki/echo-go-web-server/services"
	"github.com/labstack/echo/v4"
)

func Api(e *echo.Echo) {
	e.GET("/api/get-temperature", func(c echo.Context) error {
		temp := services.RandomWeather()
		return c.String(http.StatusOK, fmt.Sprintf("%v °C", temp))
	})
}
