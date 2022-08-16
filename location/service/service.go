package service

import (
	"context"

	"github.com/smoothbronco/community-station-api/location/pb"
	"github.com/smoothbronco/community-station-api/location/repository"
)

type Service interface {
	CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error)
	ReadLocation(ctx context.Context, req *pb.ReadLocationRequest) (*pb.ReadLocationResponse, error)
	UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.UpdateLocationResponse, error)
	DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationResponse, error)
	ListLocation(req *pb.ListLocationRequest, stream pb.LocationService_ListLocationServer) error
}

type service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{r}
}

// CreateLocation implements Service
func (s *service) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error) {
	input := req.GetLocationInput()

	id, err := s.repository.InsertLocation(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateLocationResponse{
		Location: &pb.Location{
			Id:        id,
			Name:      input.Name,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Note:      input.Note,
			MapUrl:    input.MapUrl,
		},
	}, nil
}

// DeleteLocation implements Service
func (s *service) DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationResponse, error) {
	input := req.GetId()

	err := s.repository.DeleteLocation(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteLocationResponse{Id: input}, nil
}

// ListLocation implements Service
func (s *service) ListLocation(req *pb.ListLocationRequest, stream pb.LocationService_ListLocationServer) error {
	locations, err := s.repository.SelectAllLocations()
	if err != nil {
		return err
	}
	defer locations.Close()

	for locations.Next() {
		var location pb.Location
		err := locations.Scan(&location.Id, &location.Name, &location.Latitude, &location.Longitude, &location.Note, &location.MapUrl)
		if err != nil {
			return err
		}
		stream.Send(&pb.ListLocationResponse{Location: &location})
	}

	return nil
}

// ReadLocation implements Service
func (s *service) ReadLocation(ctx context.Context, req *pb.ReadLocationRequest) (*pb.ReadLocationResponse, error) {
	input := req.GetId()
	location, err := s.repository.SelectLocationByID(ctx, input)
	if err != nil {
		return nil, err
	}
	return &pb.ReadLocationResponse{
		Location: &pb.Location{
			Id:        input,
			Name:      location.Name,
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
			Note:      location.Note,
			MapUrl:    location.MapUrl,
		},
	}, nil
}

// UpdateLocation implements Service
func (s *service) UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.UpdateLocationResponse, error) {
	id := req.GetId()
	input := req.GetLocationInput()
	err := s.repository.UpdateLocation(ctx, id, input)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateLocationResponse{
		Location: &pb.Location{
			Id:        id,
			Name:      input.Name,
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			Note:      input.Note,
			MapUrl:    input.MapUrl,
		},
	}, nil
}
