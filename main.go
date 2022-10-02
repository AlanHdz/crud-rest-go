package main

import (
	"alanhedz/golang-crud/config"
	"alanhedz/golang-crud/controllers"
	"alanhedz/golang-crud/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := config.LoadAppConfig()

	database.Connect(config.ConnectionString)
	database.Migrate()

	router := mux.NewRouter().StrictSlash(true)

	RegisterUserRoutes(router)

	log.Println(fmt.Sprintf("Starting server on port %s", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), router))
}

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
