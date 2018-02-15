package application

import (
	"github.com/graphql-go/graphql"
)

// NewRootMutationSchema is build for application root mutation
func NewRootMutationSchema() *graphql.Object {
	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "User",
			Fields: graphql.Fields{},
		},
	)
	return rootMutation
}
