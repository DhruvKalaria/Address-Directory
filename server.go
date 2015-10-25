package main

import (

	"github.com/drone/routes"
    "net/http"
    "github.com/DhruvKalaria/cmpe273-Assignment2/controllers"
    "gopkg.in/mgo.v2"
)

func main() {
    mux := routes.New()

    lc := controllers.NewLocationController(getSession())

    mux.Get("/locations/:location_id", lc.GetLocation)
    mux.Post("/locations", lc.AddLocation)
    mux.Del("/locations/:location_id", lc.DeleteLocation)
    mux.Put("/locations/:location_id", lc.UpdateLocation)
    http.Handle("/", mux)
    http.ListenAndServe(":8088", nil)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://user:user@ds041144.mongolab.com:41144/test_mongo_db")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
