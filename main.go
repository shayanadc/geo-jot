package main

import (
	auth "geo-jot/auth"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var Fields = graphql.Fields{
	"health": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "checked!", nil
		},
	},
}

var DescriptionFields = graphql.Fields{
	"description": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "description...!", nil
		},
	},
}

var MutationFields = graphql.Fields{
	"token": &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"userId": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},

		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			token, _ := auth.GenerateToken(1)
			return token, nil
		},
	},
}
var RootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: Fields}
var AuthMutation = graphql.ObjectConfig{Name: "AuthMutation", Fields: MutationFields}
var EmptyQuery = graphql.ObjectConfig{Name: "EmptyQuery", Fields: DescriptionFields}

var SchemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(RootQuery)}

var SchemaAuthConfig = graphql.SchemaConfig{Query: graphql.NewObject(EmptyQuery), Mutation: graphql.NewObject(AuthMutation)}

func GetSchema(SchemaConfig graphql.SchemaConfig) *graphql.Schema {
	schema, err := graphql.NewSchema(SchemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return &schema
}

func main() {
	http.Handle("/graphql", auth.AuthMiddleware(handler.New(&handler.Config{
		Schema: GetSchema(SchemaConfig),
	})))

	http.Handle("/auth", handler.New(&handler.Config{
		Schema: GetSchema(SchemaAuthConfig),
	}))

	log.Println("geo-jot *** gql server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
