package handler

import (
	"net/http"
	"os"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
}

type Response struct {
	Users []Users `json:"users" bson:"users"`
}

type Oops struct {
	Oops string `json:"oops" bson:"oops"`
}

func handler(c echo.Context) error {
	err := mgm.SetDefaultConfig(nil, os.Getenv("MONGO_DATA"), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		oops := Oops{Oops: "connection to mongodb fails"}
		return c.JSON(http.StatusInternalServerError, oops)
	}

	users := []Users{}
	mgm.Coll(&Users{}).SimpleFind(&users, bson.M{})

	response := Response{Users: users}
	return c.JSON(http.StatusOK, response)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/users", handler)
	e.ServeHTTP(w, r)
}
