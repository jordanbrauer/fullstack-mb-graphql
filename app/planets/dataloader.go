package planets

import (
	"context"
	"fmt"
	"fullstackmb/app"
	"log"
	"strings"

	"github.com/graph-gophers/dataloader"
)

func Load(ctx context.Context, id int) (*Schema, error) {
	result, err := loader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%d", id)))()

	return result.(*Schema), err
}

var loader = dataloader.NewBatchedLoader(func(ctx context.Context, loading dataloader.Keys) []*dataloader.Result {
	results := make([]*dataloader.Result, 0)
	query, err := app.Queries.ReadFile("database/planet_batch.sql")

	if err != nil {
		log.Panic("cannot read query for data loader", err)

		return results
	}

	rows, err := app.Database().Query(
		strings.ReplaceAll(string(query), "?", strings.Join(loading.Keys(), ",")),
	)

	if err != nil {
		log.Panic(err)

		return results
	}

	planets := map[string]*Schema{}
	errors := map[string]error{}

	defer rows.Close()

	for rows.Next() {
		planet := new(Schema)
		err = rows.Scan(&planet.ID, &planet.Name)
		key := fmt.Sprintf("%d", planet.ID)
		errors[key] = err
		planets[key] = planet
	}

	log.Printf("Queried %d planet(s) for %d character(s)\n", len(planets), len(loading.Keys()))

	for _, key := range loading.Keys() {
		results = append(results, &dataloader.Result{
			Data:  planets[key],
			Error: errors[key],
		})
	}

	return results
}, dataloader.WithCache(&dataloader.NoCache{}))
