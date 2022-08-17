package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/smoothbronco/community-station-api/graph/generated"
	"github.com/smoothbronco/community-station-api/graph/model"
)

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input model.CreateInput) (*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, input model.UpdateInput) (*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteLocation is the resolver for the deleteLocation field.
func (r *mutationResolver) DeleteLocation(ctx context.Context, input int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, input int) (*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
