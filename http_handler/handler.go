package http_handler

import (
	"geo-jot/routes"
	"net/http"
	"net/http/httptest"

	"github.com/graphql-go/handler"
)

func Router(routes []routes.Route) {

	for _, route := range routes {

		handler := RouteHandler(route)

		http.Handle(route.Name, handler)
	}
}

func RouteHandler(route routes.Route) http.Handler {

	schemConfig := handler.Config{
		Schema: route.Handler,
	}

	handler := Handler(&schemConfig)

	if route.Middleware != nil {

		handler = route.Middleware(Handler(&schemConfig))
	}

	return handler
}

func Handler(handlerConfig *handler.Config) http.Handler {

	return handler.New(handlerConfig)
}

func HandlerRecorder(name string, req *http.Request) *httptest.ResponseRecorder {

	recorder := httptest.NewRecorder()

	for _, route := range routes.Routes {

		if name == route.Name {

			gqlHandler := RouteHandler(route)

			gqlHandler.ServeHTTP(recorder, req)

		}
	}
	return recorder
}
