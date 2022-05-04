package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"fullstackmb/app/characters"
	"fullstackmb/app/films"
	"fullstackmb/graph/generated"
	"fullstackmb/graph/model"
)

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

// Film returns generated.FilmResolver implementation.
func (r *Resolver) Film() generated.FilmResolver { return &filmResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type filmResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *filmResolver) ID(ctx context.Context, obj *model.Film) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
