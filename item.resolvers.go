package main

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/breadchris/hn/ent"
)

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input ent.CreateItemInput) (*ent.Item, error) {
	return r.client.Item.Create().
		SetID("test").
		SetDeleted(input.Deleted).
		SetDead(input.Dead).
		SetType(input.Type).
		SetTime(input.Time).
		Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
