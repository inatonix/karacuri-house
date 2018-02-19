package application

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

var userGQLobject = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
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
},
)

// NewRootMutationSchema is build for application root mutation
func NewRootMutationSchema() *graphql.Object {
	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootMutation",
			Fields: graphql.Fields{
				"createUser": &graphql.Field{
					Type:        userGQLobject,
					Description: "Create New User",
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"age": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						name, _ := params.Args["name"].(string)
						age, _ := params.Args["age"].(int)
						newID := 10

						newUser := User{
							ID:   newID,
							Name: name,
							Age:  age,
						}

						return newUser, nil
					},
				},
			},
		},
	)

	return rootMutation
}
