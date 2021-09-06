package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type version struct {
	Version string `json:"version"`
}

type Response struct {
	Version string `json:"version" bson:"version"`
}

type Oops struct {
	Oops string `json:"oops" bson:"oops"`
}

func handler(c echo.Context) error {
	response := Response{Version: "0.0.0"}
	return c.JSON(http.StatusOK, response)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/", handler)
	e.ServeHTTP(w, r)
}
