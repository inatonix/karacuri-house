package application

import (
	"github.com/graphql-go/graphql"
)

// NewRootQuerySchema is build for application root query
func NewRootQuerySchema() *graphql.Object {
	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: graphql.Fields{},
		},
	)

	return rootQuery
}
