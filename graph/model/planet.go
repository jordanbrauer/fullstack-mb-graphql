package model

import "fullstackmb/app/planets"

type Planet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (planet *Planet) Fill(from *planets.Schema) *Planet {
	planet.ID = from.ID
	planet.Name = from.Name

	return planet
}
