package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/karacuri-house/application"
)

func main() {

	tApplicationRouter := chi.NewRouter()
	tApplicationRouter.Use(middleware.Logger)
	tApplicationRouter.Get("/graphiql", application.GraphQLRequest)
	tApplicationRouter.Post("/graphiql", application.GraphQLRequest)
	log.Printf("trace: this is a trace log.")
	http.ListenAndServe(":8080", tApplicationRouter)

}
