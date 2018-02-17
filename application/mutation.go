package application

import (
	"github.com/graphql-go/graphql"
)

// NewRootMutationSchema is build for application root mutation
func NewRootMutationSchema() *graphql.Object {
	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
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
	return rootMutation
}
