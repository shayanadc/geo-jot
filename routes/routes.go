package routes

import (
	"geo-jot/api_schema"
	"geo-jot/auth"
	"net/http"

	"github.com/graphql-go/graphql"
)

type Route struct {
	Name       string
	Handler    *graphql.Schema
	Middleware func(http.Handler) http.Handler
}

var Routes = []Route{
	{
		Name:       "/graphql",
		Handler:    api_schema.CreateSchema(api_schema.SchemaConfig),
		Middleware: auth.AuthMiddleware,
	},
	{
		Name:    "/auth",
		Handler: api_schema.CreateSchema(api_schema.SchemaAuthConfig),
	},
}
