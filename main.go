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

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/ini.v1"
)

var server *handler.Server
var config *ini.File

//go:embed database/*.sql
var queries embed.FS

func main() {
	graphql := config.Section("graphql")
	host := graphql.Key("host").Value()
	port, _ := graphql.Key("port").Int()
	path := graphql.Key("path").Value()

	http.Handle(fmt.Sprintf("/%s", graphql.Key("path").Value()), server)
	log.Printf("Server listening on %s:%d", host, port)
	log.Printf("GraphQL API available at http://%s:%d/%s", host, port, path)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}

func init() {
	config = app.Config("./config/app.ini")

	app.Database()
	app.UseQueries(queries)

	server = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			Config: config,
		},
	}))
}
