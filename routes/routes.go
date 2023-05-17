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
	r.HandleFunc("/", controllers.Home)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
