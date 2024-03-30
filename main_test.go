package main

import (
	"bytes"
	"encoding/json"
	auth "geo-jot/auth"
	graphql "geo-jot/graphql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/graphql-go/handler"
)

func TestGraphQLHealthEndpoint(t *testing.T) {

	token, _ := auth.GenerateToken(1)

	query := `{"query":"{ health }"}`

	req, err := http.NewRequest("POST", "/graphql", bytes.NewBufferString(query))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", token)

	recorder := httptest.NewRecorder()

	gqlHandler := auth.AuthMiddleware(handler.New(&handler.Config{Schema: graphql.GetSchema(graphql.SchemaConfig)}))

	gqlHandler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedBody := `{"data":{"health":"checked!"}}`
	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body, expectedBody)
	}
}

func TestCreateUserEndpoint(t *testing.T) {

	token, _ := auth.GenerateToken(1)

	query := `mutation {
		createUser(name: "John Doe", location: {lat: 51.5, lon: -0.12}) {
			name,
			location{
				lat,
				lon
			}
		}
	}`

	reqBody, _ := json.Marshal(map[string]string{
		"query": query,
	})

	req, err := http.NewRequest("POST", "/graphql", bytes.NewBufferString(string(reqBody)))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", token)

	recorder := httptest.NewRecorder()

	gqlHandler := auth.AuthMiddleware(handler.New(&handler.Config{Schema: graphql.GetSchema(graphql.SchemaConfig)}))

	gqlHandler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedBody := `{"data":{"createUser":{"location":{"lat":51.5,"lon":-0.12},"name":"John Doe"}}}`
	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body, expectedBody)
	}
}
