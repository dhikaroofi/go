package main

import (
	"github.com/dhikaroofi/go/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	route := mux.NewRouter()
	app := app.Mojoo{}
	app.Initialize(route)
	log.Print("App is running on ")
	log.Fatal(http.ListenAndServe(":8989", route))

}
