package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/smoothbronco/community-station-api/graph/generated"
	"github.com/smoothbronco/community-station-api/graph/model"
	"github.com/smoothbronco/community-station-api/location/pb"
)

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input model.CreateInput) (*model.Location, error) {
	location, err := r.LocationClient.CreateLocation(
		ctx,
		&pb.LocationInput{
			Name:      input.Name,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Note:      input.Note,
			MapUrl:    input.MapUrl,
		})
	if err != nil {
		return nil, err
	}

	return location, nil
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, input model.UpdateInput) (*model.Location, error) {
	location, err := r.LocationClient.UpdateLocation(
		ctx,
		int64(input.ID),
		&pb.LocationInput{
			Name:      input.Name,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Note:      input.Note,
			MapUrl:    input.MapUrl,
		})
	if err != nil {
		return nil, err
	}

	return location, nil
}

// DeleteLocation is the resolver for the deleteLocation field.
func (r *mutationResolver) DeleteLocation(ctx context.Context, input int) (int, error) {
	id, err := r.LocationClient.DeleteLocation(ctx, int64(input))
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, input int) (*model.Location, error) {
	location, err := r.LocationClient.ReadLocation(ctx, int64(input))
	if err != nil {
		return nil, err
	}

	return location, nil
}

// Locations is the resolver for the locations field.
func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	locations, err := r.LocationClient.ListLocation(ctx)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
