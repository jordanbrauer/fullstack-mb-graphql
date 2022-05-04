package planets

import (
	"fullstackmb/app"
	"log"
)

type Schema struct {
	ID   int
	Name string
}

// Find a single planet by it's unique ID.
func Find(id int) (*Schema, error) {
	var planet *Schema

	query, err := app.Queries.ReadFile("database/planet.sql")

	if err != nil {
		return planet, err
	}

	rows, err := app.Database().Query(string(query), id)

	if err != nil {
		return planet, err
	}

	defer rows.Close()

	for rows.Next() {
		planet = new(Schema)
		err = rows.Scan(&planet.ID, &planet.Name)

		if err != nil {
			return planet, err
		}
	}

	log.Printf("Queried planet %d,%s", planet.ID, planet.Name)

	return planet, nil
}
