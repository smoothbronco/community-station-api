package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/smoothbronco/community-station-api/graph/model"
	"github.com/smoothbronco/community-station-api/location/pb"
	"google.golang.org/grpc"
)

func TestNewClient(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Close(t *testing.T) {
	type fields struct {
		conn    *grpc.ClientConn
		Service pb.LocationServiceClient
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:    tt.fields.conn,
				Service: tt.fields.Service,
			}
			c.Close()
		})
	}
}

func TestClient_CreateLocation(t *testing.T) {
	type fields struct {
		conn    *grpc.ClientConn
		Service pb.LocationServiceClient
	}
	type args struct {
		ctx   context.Context
		input *pb.LocationInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:    tt.fields.conn,
				Service: tt.fields.Service,
			}
			got, err := c.CreateLocation(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ReadLocation(t *testing.T) {
	type fields struct {
		conn    *grpc.ClientConn
		Service pb.LocationServiceClient
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:    tt.fields.conn,
				Service: tt.fields.Service,
			}
			got, err := c.ReadLocation(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ReadLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ReadLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_UpdateLocation(t *testing.T) {
	type fields struct {
		conn    *grpc.ClientConn
		Service pb.LocationServiceClient
	}
	type args struct {
		ctx   context.Context
		id    int64
		input *pb.LocationInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:    tt.fields.conn,
				Service: tt.fields.Service,
			}
			got, err := c.UpdateLocation(tt.args.ctx, tt.args.id, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdateLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.UpdateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DeleteLocation(t *testing.T) {
	type fields struct {
		conn    *grpc.ClientConn
		Service pb.LocationServiceClient
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:    tt.fields.conn,
				Service: tt.fields.Service,
			}
			got, err := c.DeleteLocation(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.DeleteLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_ListLocation(t *testing.T) {
	type fields struct {
		conn    *grpc.ClientConn
		Service pb.LocationServiceClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:    tt.fields.conn,
				Service: tt.fields.Service,
			}
			got, err := c.ListLocation(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.ListLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
