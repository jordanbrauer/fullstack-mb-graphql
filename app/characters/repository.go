package characters

import (
	"database/sql"
	"fmt"
	"fullstackmb/app"
)

type Schema struct {
	ID, HomeworldId                                                                        int
	Name, Created, Edited, HairColor, SkinColor, EyeColor, BirthYear, Gender, Mass, Height string
}

// Query for all known characters
func All() ([]*Schema, error) {
	rows, err := query("characters")

	if err != nil {
		return []*Schema{}, err
	}

	return scan(rows)
}

// Query for characters in the given film
func AppearingIn(film int) ([]*Schema, error) {
	rows, err := query("film_characters", film)

	if err != nil {
		return []*Schema{}, err
	}

	return scan(rows)
}

func query(name string, parameters ...any) (*sql.Rows, error) {
	query, err := app.Queries.ReadFile(fmt.Sprintf("database/%s.sql", name))

	if err != nil {
		return nil, err
	}

	return app.Database().Query(string(query), parameters...)
}

func scan(rows *sql.Rows) ([]*Schema, error) {
	defer rows.Close()

	characters := []*Schema{}

	for rows.Next() {
		character := &Schema{}
		err := rows.Scan(
			&character.ID,
			&character.Created,
			&character.Edited,
			&character.Name,
			&character.Height,
			&character.Mass,
			&character.HairColor,
			&character.SkinColor,
			&character.EyeColor,
			&character.BirthYear,
			&character.Gender,
			&character.HomeworldId,
		)

		if err != nil {
			return characters, err
		}

		characters = append(characters, character)
	}

	if err := rows.Err(); err != nil {
		return characters, err
	}

	return characters, nil
}
