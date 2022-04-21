package routes

import (
	"golang-crud/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RoutingHandler() {
	router := mux.NewRouter()

	//book end points
	router.HandleFunc("/api/v1/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/v1/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/v1/book/{bid}", controllers.GetSingleBook).Methods("GET")
	router.HandleFunc("/api/v1/book/{bid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/v1/book/{bid}", controllers.Deletebook).Methods("DELETE")
	router.HandleFunc("/api/v1/user", controllers.IsAuthenticated(controllers.UserHome)).Methods("GET")
	router.HandleFunc("/api/v1/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/v1/login", controllers.Login).Methods("POST")

	// Serve Sever
	log.Println("Listening on port :4000 -> http://localhost:4000")
	http.ListenAndServe(":4000", router)
}
