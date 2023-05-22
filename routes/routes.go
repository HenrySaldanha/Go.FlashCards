package routes

import (
	"log"
	"net/http"

	"github.com/HenrySaldanha/Go.FlashCards/controllers"
	"github.com/HenrySaldanha/Go.FlashCards/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/cards", controllers.GetAllCards).Methods("GET")
	r.HandleFunc("/cards/{id}", controllers.GetCardById).Methods("GET")
	r.HandleFunc("/cards/{id}", controllers.DeleteCard).Methods("DELETE")
	r.HandleFunc("/cards", controllers.InsertCard).Methods("POST")
	r.HandleFunc("/cards/{id}", controllers.UpdateCard).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
