package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/smoothbronco/community-station-api/graph"
	"github.com/smoothbronco/community-station-api/graph/generated"
	"github.com/smoothbronco/community-station-api/location/client"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	locationClient, err := client.NewClient("localhost:50000")
	if err != nil {
		locationClient.Close()
		log.Fatalf("Failed to create location client: %v\n", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{LocationClient: locationClient}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}