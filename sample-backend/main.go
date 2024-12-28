package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	restUrl     = "http://sample-restapi.ecs.internal:8080"
	restApiPath = "/api/v1/date"
	host        = "0.0.0.0"
	port        = "8070"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "health check")
	})

	e.GET("/api", func(c echo.Context) error {
		resp, err := http.Get(restUrl + restApiPath)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to get date")
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to read response body")
		}

		return c.String(resp.StatusCode, string(body))
	})

	e.Start(host + ":" + port)
}
