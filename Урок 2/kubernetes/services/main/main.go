package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const AppServiceHost = "http://app"

// поход в сервис app
func RequestExternalService() int {
	resp, err := http.Get(AppServiceHost + "/hello")
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError
	} else {
		log.Println("OK")
	}

	return resp.StatusCode
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(RequestExternalService(), "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
