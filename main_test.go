package main

import (
	"bytes"
	"fmt"
	auth "geo-jot/Auth"
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

	gqlHandler := auth.AuthMiddleware(handler.New(&handler.Config{Schema: GetSchema(SchemaConfig)}))

	gqlHandler.ServeHTTP(recorder, req)

	fmt.Println(recorder.Body.String())

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expectedBody := `{"data":{"health":"checked!"}}`
	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v, want %v", body, expectedBody)
	}
}
