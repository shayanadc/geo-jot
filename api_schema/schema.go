package api_schema

import (
	"geo-jot/auth"
	"geo-jot/models"
	"log"

	"github.com/graphql-go/graphql"
)

var Fields = graphql.Fields{
	"health": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "checked!", nil
		},
	},
}

var LocationInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "LocationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"lat": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
		"lon": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Float),
		},
	},
})
var RootMutationFields = graphql.Fields{
	"createUser": &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"location": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(LocationInputType),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {

			name, _ := p.Args["name"].(string)

			location, _ := p.Args["location"].(map[string]interface{})

			lat, _ := location["lat"].(float64)
			lon, _ := location["lon"].(float64)

			return models.CreateUser(name, models.Location{Lat: lat, Lon: lon}), nil
		},
	},
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: LocationType,
		},
	},
})

var LocationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Fields: graphql.Fields{
		"lat": &graphql.Field{
			Type: graphql.Float,
		},
		"lon": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

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
var RootMutation = graphql.ObjectConfig{Name: "RootMutation", Fields: RootMutationFields}
var AuthMutation = graphql.ObjectConfig{Name: "AuthMutation", Fields: MutationFields}
var EmptyQuery = graphql.ObjectConfig{Name: "EmptyQuery", Fields: DescriptionFields}

var SchemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(RootQuery), Mutation: graphql.NewObject(RootMutation)}
var SchemaAuthConfig = graphql.SchemaConfig{Query: graphql.NewObject(EmptyQuery), Mutation: graphql.NewObject(AuthMutation)}

func CreateSchema(SchemaConfig graphql.SchemaConfig) *graphql.Schema {
	schema, err := graphql.NewSchema(SchemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return &schema
}
