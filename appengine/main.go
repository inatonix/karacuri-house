package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/karacuri-house/application"
)

func main() {

	log.Println("Get single product: curl -g 'http://localhost:8080/graphql?query={product(id:1){id,title,url,image,likes_count,access_token}}'")

	tApplicationRouter := chi.NewRouter()
	tApplicationRouter.Use(middleware.Logger)
	tApplicationRouter.Get("/graphiql", application.GraphQLRequest)
	tApplicationRouter.Post("/graphiql", application.GraphQLRequest)

	http.ListenAndServe(":8080", tApplicationRouter)

}
