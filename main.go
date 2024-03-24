package main

import (
	auth "geo-jot/auth"
	graphql "geo-jot/graphql"

	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	http.Handle("/graphql", auth.AuthMiddleware(handler.New(&handler.Config{
		Schema: graphql.GetSchema(graphql.SchemaConfig),
	})))

	http.Handle("/auth", handler.New(&handler.Config{
		Schema: graphql.GetSchema(graphql.SchemaAuthConfig),
	}))

	log.Println("geo-jot *** gql server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
