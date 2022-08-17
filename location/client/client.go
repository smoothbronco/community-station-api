package client

import (
	"context"
	"io"

	"github.com/smoothbronco/community-station-api/graph/model"
	"github.com/smoothbronco/community-station-api/location/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	Service pb.LocationServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewLocationServiceClient(conn)

	return &Client{conn, client}, nil
}

func (c *Client) CreateLocation(ctx context.Context, input *pb.LocationInput) (*model.Location, error) {
	res, err := c.Service.CreateLocation(
		ctx,
		&pb.CreateLocationRequest{LocationInput: input},
	)
	if err != nil {
		return nil, err
	}

	return &model.Location{
		ID:        int(res.Location.Id),
		Name:      res.Location.Name,
		Latitude:  res.Location.Latitude,
		Longitude: res.Location.Longitude,
		Note:      res.Location.Note,
	}, nil
}

func (c *Client) ReadLocation(ctx context.Context, id int64) (*model.Location, error) {
	res, err := c.Service.ReadLocation(ctx, &pb.ReadLocationRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &model.Location{
		ID:        int(res.Location.Id),
		Name:      res.Location.Name,
		Latitude:  res.Location.Latitude,
		Longitude: res.Location.Longitude,
		Note:      res.Location.Note,
	}, nil
}

func (c *Client) UpdateLocation(ctx context.Context, id int64, input *pb.LocationInput) (*model.Location, error) {
	res, err := c.Service.UpdateLocation(ctx, &pb.UpdateLocationRequest{Id: id, LocationInput: input})
	if err != nil {
		return nil, err
	}

	return &model.Location{
		ID:        int(res.Location.Id),
		Name:      res.Location.Name,
		Latitude:  res.Location.Latitude,
		Longitude: res.Location.Longitude,
		Note:      res.Location.Note,
	}, nil
}

func (c *Client) DeleteLocation(ctx context.Context, id int64) (int64, error) {
	res, err := c.Service.DeleteLocation(ctx, &pb.DeleteLocationRequest{Id: id})
	if err != nil {
		return 0, err
	}
	return res.Id, nil
}

func (c *Client) ListLocation(ctx context.Context) ([]*model.Location, error) {
	res, err := c.Service.ListLocation(ctx, &pb.ListLocationRequest{})
	if err != nil {
		return nil, err
	}

	var locations []*model.Location
	for {
		r, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		locations = append(locations, &model.Location{
			ID:        int(r.Location.Id),
			Name:      r.Location.Name,
			Latitude:  r.Location.Latitude,
			Longitude: r.Location.Longitude,
			Note:      r.Location.Note,
		})
	}
	return locations, nil
}
