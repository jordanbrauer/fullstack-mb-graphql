package model

import (
	"fullstackmb/app/films"
	"strings"
)

type Film struct {
	ID int `json:"id"`
	// The film's theatric title.
	Title string `json:"title"`
	// The chronological order of the film in relation to all the others.
	Episode int `json:"episode"`
	// A series of one or more paragraphs which are scrolled across space and time
	// in big yellow letters.
	Crawl string `json:"crawl"`
	// Name of the director in charge of the movie.
	Director string `json:"director"`
	// A list of producers that worked on the film.
	Producers []string `json:"producers"`
	// All known characters that appear in the film.
	Characters []*Character `json:"characters"`
	// Theatric release date of the movie.
	ReleasedAt string `json:"releasedAt"`
}

func (film *Film) Fill(from *films.Schema) *Film {
	producers := strings.Split(from.Producer, ",")

	for index, producer := range producers {
		producers[index] = strings.Trim(producer, " ")
	}

	film.ID = from.ID
	film.Episode = from.EpisodeID
	film.Title = from.Title
	film.Crawl = from.OpeningCrawl
	film.Director = from.Director
	film.Producers = producers
	film.ReleasedAt = from.ReleaseDate

	return film
}
