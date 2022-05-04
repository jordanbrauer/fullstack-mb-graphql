package films

import (
	"fullstackmb/app"
)

type Schema struct {
	ID, EpisodeID                                        int
	Title, OpeningCrawl, Director, Producer, ReleaseDate string
}

func All() ([]*Schema, error) {
	films := []*Schema{}
	query, err := app.Queries.ReadFile("database/films.sql")

	if err != nil {
		return films, err
	}

	rows, err := app.Database().Query(string(query))

	if err != nil {
		return films, err
	}

	defer rows.Close()

	for rows.Next() {
		var created, edited string

		film := &Schema{}
		err = rows.Scan(
			&film.ID,
			&film.ReleaseDate,
			&created,
			&edited,
			&film.Title,
			&film.EpisodeID,
			&film.OpeningCrawl,
			&film.Director,
			&film.Producer,
		)

		if err != nil {
			return films, err
		}

		films = append(films, film)
	}

	return films, nil
}
