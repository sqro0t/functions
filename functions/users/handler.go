package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/kamva/mgm/v3"
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

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := mgm.SetDefaultConfig(nil, os.Getenv("MONGO_DATA"), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		oops := Oops{Oops: "connection to mongodb fails"}
		json.NewEncoder(w).Encode(oops)
	}

	users := []Users{}
	mgm.Coll(&Users{}).SimpleFind(&users, bson.M{})

	response := Response{Users: users}
	json.NewEncoder(w).Encode(response)
}
