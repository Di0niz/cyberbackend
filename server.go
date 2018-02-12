package main

import (
	"net/http"

	"github.com/Di0niz/cyberbackend/config"
	"github.com/Di0niz/cyberbackend/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
	//	"github.com/prometheus/client_golang/prometheus/promhttp"
	//"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handlers.Handler{
		DB: config.MySQLConnect(),
	}

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Helo, World!\n")
	})

	api_ver := "/api/v1"

	e.GET(api_ver+"/user", h.GetUser)

	// получаем информацию по командам
	e.GET(api_ver+"/team/:id", h.GetTeam)
	e.GET(api_ver+"/team/list", h.GetListTeam)
	e.POST(api_ver+"/team", h.PostTeam)
	e.PUT(api_ver+"/team/:id", h.PutTeam)
	e.DELETE(api_ver+"/team/:id", h.DeleteTeam)

	//e.GET("/metrics", promhttp.Handler())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
