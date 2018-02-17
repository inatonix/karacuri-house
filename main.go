package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/karacuri-house/application"
	"github.com/karacuri-house/infrastructure/database"
)

func main() {
	log.Printf("root:" + infrastructure.PASSWORD + "@unix(/cloudsql/" + infrastructure.PROJECT_NAME + ":us-central1:karacuri-house/karacuri-house?charset=utf8&parseTime=True")

	tApplicationRouter := chi.NewRouter()
	tApplicationRouter.Use(middleware.Logger)
	tApplicationRouter.Get("/graphiql", application.GraphQLRequest)
	tApplicationRouter.Post("/graphiql", application.GraphQLRequest)
	log.Printf("trace: this is a trace log.")
	http.ListenAndServe(":8080", tApplicationRouter)

}
