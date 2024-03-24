package main

import (
	"bytes"
	"context"
	auth "geo-jot/auth"
	graphql "geo-jot/graphql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	"github.com/graphql-go/handler"
)

var recorder *httptest.ResponseRecorder
var token string

func iAmAnAuthenticatedUser() (err error) {
	token, err = auth.GenerateToken(1)
	if err != nil {
		return err
	}
	return nil
}
func iSendRequestToWithBody(method, endpoint string, body *godog.DocString) (err error) {

	req, err := http.NewRequest("POST", "/graphql", bytes.NewBufferString(body.Content))

	if err != nil {
		return err
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	gqlHandler := auth.AuthMiddleware(handler.New(&handler.Config{Schema: graphql.GetSchema(graphql.SchemaConfig)}))

	gqlHandler.ServeHTTP(recorder, req)

	return nil
}

func iShouldRecieve(status int) (err error) {
	if recorder.Code != status {
		return godog.ErrPending
	}
	return nil
}

func TestFeatures(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		recorder = httptest.NewRecorder()
		return ctx, nil
	})
	ctx.Step(`^I am an authenticated user$`, iAmAnAuthenticatedUser)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)" with body:$`, iSendRequestToWithBody)
	ctx.Step(`^I should recieve "([^"]*)"$`, iShouldRecieve)
}
