package main

import (
	"geo-jot/http_handler"
	"geo-jot/routes"
	"log"
	"net/http"
)

func main() {

	App()

	log.Println("geo-jot *** gql server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func App() {

	http_handler.Router(routes.Routes)
}
