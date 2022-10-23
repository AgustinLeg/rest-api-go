package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/agustinleg/postgres-rest-api/db"
	"github.com/agustinleg/postgres-rest-api/models"
	"github.com/agustinleg/postgres-rest-api/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
