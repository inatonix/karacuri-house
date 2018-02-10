package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/karacuri-house/app"
)

func main() {

	log.Println("LocalServer is Running at " + strconv.Itoa(aConfig.Port) + " Port ... ")
	log.Println("Get single product: curl -g 'http://localhost:8080/graphql?query={product(id:1){id,title,url,image,likes_count,access_token}}'")

	tApplicationRouter := chi.NewRouter()
	tApplicationRouter.Use(middleware.Logger)
	tApplicationRouter.Get("/graphql", app.Request)
	tApplicationRouter.Get("/graphiql", app.GraphiQLRequest)
	tApplicationRouter.Post("/graphiql", app.GraphiQLRequest)

	return http.ListenAndServe(fmt.Sprintf(":%d", aConfig.Port), tApplicationRouter)

}
