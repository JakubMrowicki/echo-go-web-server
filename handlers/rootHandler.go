package handlers

import (
	"context"
	"fmt"

	"github.com/JakubMrowicki/echo-go-web-server/components"
	"github.com/labstack/echo/v4"
)

func Headers(c echo.Context) error {
	req := []string{}
	for key, values := range c.Request().Header {
		line := ""
		line += fmt.Sprint(key + ": ")
		for _, value := range values {
			line += fmt.Sprintln(value)
		}
		req = append(req, line)
	}
	html := components.Home(req)
	return html.Render(context.Background(), c.Response().Writer)
}
