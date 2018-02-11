package application

import (
	"net/http"

	"github.com/graphql-go/graphql"
	graphiql "github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    NewRootQuerySchema(),
		Mutation: NewRootMutationSchema(),
	},
)

func GraphQLRequest(w http.ResponseWriter, r *http.Request) {
	handler := graphiql.New(
		&graphiql.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		},
	)
	handler.ServeHTTP(w, r)
}
