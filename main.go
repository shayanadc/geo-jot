package main

import (
	"fmt"
	auth "geo-jot/Auth"
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
var RootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: Fields}
var SchemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(RootQuery)}

func GetSchama() *graphql.Schema {
	schema, err := graphql.NewSchema(SchemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return &schema
}

func main() {
	fmt.Println(auth.GenerateToken(1))
	http.Handle("/graphql", auth.AuthMiddleware(handler.New(&handler.Config{
		Schema: GetSchama(),
	})))

	log.Println("geo-jot *** gql server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
