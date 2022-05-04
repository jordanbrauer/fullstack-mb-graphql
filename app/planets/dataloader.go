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

var loader = dataloader.NewBatchedLoader(func(ctx context.Context, k dataloader.Keys) []*dataloader.Result {
	query, err := app.Queries.ReadFile("database/planet_batch.sql")

	if err != nil {
		log.Panic("cannot read query for data loader", err)
	}

	results := make([]*dataloader.Result, 0)
	ids := strings.Join(k.Keys(), ",")
	body := strings.ReplaceAll(string(query), "?", ids)
	rows, err := app.Database().Query(body)

	log.Printf("Querying planets for %d character(s)\n", len(k.Keys()))

	if err != nil {
		log.Panic(err)

		return results
	}

	m := map[string]*Schema{}

	defer rows.Close()

	for rows.Next() {
		result := new(Schema)
		err = rows.Scan(&result.ID, &result.Name)

		if err != nil {
			log.Panic(err.Error())
		}

		m[fmt.Sprintf("%d", result.ID)] = result
	}

	for _, key := range k.Keys() {
		results = append(results, &dataloader.Result{Data: m[key]})
	}

	return results
}, dataloader.WithCache(&dataloader.NoCache{}))
