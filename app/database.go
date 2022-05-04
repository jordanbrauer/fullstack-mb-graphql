package app

import (
	"database/sql"
	"embed"
	"log"
)

var database *sql.DB
var Queries embed.FS

func Database() *sql.DB {
	if nil != database {
		return database
	}

	var err error

	database, err = sql.Open("sqlite3", "./database/swapi.db?cache=shared&mode=memory")

	if err != nil {
		log.Panic("unable to connect to database!", err)
	}

	return database
}

func UseQueries(filesystem embed.FS) {
	Queries = filesystem
}
