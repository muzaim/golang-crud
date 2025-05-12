package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/routes"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running on Port : ", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
	log.Println("dasadsa")
}

// video
// https://www.youtube.com/watch?v=Ywv0QUjUfCE&t=146s
// menit ke 35:01
