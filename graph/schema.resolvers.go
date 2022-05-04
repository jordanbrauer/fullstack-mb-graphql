package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fullstackmb/app/characters"
	"fullstackmb/app/films"
	"fullstackmb/app/planets"
	"fullstackmb/graph/generated"
	"fullstackmb/graph/model"
)

func (r *characterResolver) Homeworld(ctx context.Context, obj *model.Character) (*model.Planet, error) {
	planet, err := planets.Find(obj.HomeworldID)
	// planet, err := planets.Load(ctx, obj.HomeworldID)

	return new(model.Planet).Fill(planet), err
}

func (r *filmResolver) Characters(ctx context.Context, obj *model.Film) ([]*model.Character, error) {
	nodes := []*model.Character{}
	appeared, err := characters.AppearingIn(obj.ID)

	if err != nil {
		return nodes, err
	}

	for _, appearance := range appeared {
		nodes = append(nodes, new(model.Character).Fill(appearance))
	}

	return nodes, nil
}

func (r *mutationResolver) Sum(ctx context.Context, numbers []int) (int, error) {
	var sum int

	for _, operand := range numbers {
		sum += operand
	}

	return sum, nil
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func (r *queryResolver) Server(ctx context.Context) (*model.Info, error) {
	config := r.Config.Section("service")

	return &model.Info{
		Name:    config.Key("name").Value(),
		Version: config.Key("version").Value(),
		Running: true,
	}, nil
}

func (r *queryResolver) Films(ctx context.Context) ([]*model.Film, error) {
	nodes := []*model.Film{}
	movies, err := films.All()

	if err != nil {
		return nodes, err
	}

	for _, movie := range movies {
		nodes = append(nodes, new(model.Film).Fill(movie))
	}

	return nodes, nil
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	all, err := characters.All()
	nodes := []*model.Character{}

	if err != nil {
		return nodes, err
	}

	for _, appearance := range all {
		nodes = append(nodes, new(model.Character).Fill(appearance))
	}

	return nodes, nil
}

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

// Film returns generated.FilmResolver implementation.
func (r *Resolver) Film() generated.FilmResolver { return &filmResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type characterResolver struct{ *Resolver }
type filmResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
