package main

//go:generate go run github.com/99designs/gqlgen

import (
	"embed"
	"fmt"
	"fullstackmb/app"
	"fullstackmb/graph"
	"fullstackmb/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
)

var host, port, path string
var server *handler.Server

//go:embed database/*.sql
var queries embed.FS

func main() {
	http.Handle(fmt.Sprintf("/%s", path), server)
	log.Printf("Server listening on %s:%s", host, port)
	log.Printf("GraphQL API available at http://%s:%s/%s", host, port, path)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil))
}

func init() {
	host = os.Getenv("FSMB_HOST")
	port = os.Getenv("FSMB_PORT")
	path = os.Getenv("FSMB_PATH")

	if "" == host {
		host = "localhost"
	}

	if "" == port {
		port = "1337"
	}

	if "" == path {
		path = "graphql"
	}

	app.Database()
	app.UseQueries(queries)

	server = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))
}
