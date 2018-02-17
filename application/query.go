package application

import (
	"github.com/graphql-go/graphql"
)

// NewRootQuerySchema is build for application root query
func NewRootQuerySchema() *graphql.Object {
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: graphql.NewObject(graphql.ObjectConfig{
						Name: "Params",
						Fields: graphql.Fields{
							"id": &graphql.Field{
								Type: graphql.Int,
							},
							"name": &graphql.Field{
								Type: graphql.String,
							},
							"age": &graphql.Field{
								Type: graphql.Int,
							},
						},
					}),
				},
			},
		},
	)

	return rootQuery
}
