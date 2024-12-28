package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	host = "0.0.0.0"
	port = "8080"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.GET("/api/v1/date", func(c echo.Context) error {
		currentTime := time.Now().Format(time.RFC3339)
		return c.String(http.StatusOK, currentTime)
	})

	e.Start(host + ":" + port)
}
