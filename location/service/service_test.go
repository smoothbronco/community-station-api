package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/smoothbronco/community-station-api/location/pb"
	"github.com/smoothbronco/community-station-api/location/repository"
)

func TestNewService(t *testing.T) {
	type args struct {
		r repository.Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateLocation(t *testing.T) {
	type fields struct {
		repository repository.Repository
	}
	type args struct {
		ctx context.Context
		req *pb.CreateLocationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateLocationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repository: tt.fields.repository,
			}
			got, err := s.CreateLocation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteLocation(t *testing.T) {
	type fields struct {
		repository repository.Repository
	}
	type args struct {
		ctx context.Context
		req *pb.DeleteLocationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteLocationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repository: tt.fields.repository,
			}
			got, err := s.DeleteLocation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.DeleteLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ListLocation(t *testing.T) {
	type fields struct {
		repository repository.Repository
	}
	type args struct {
		req    *pb.ListLocationRequest
		stream pb.LocationService_ListLocationServer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repository: tt.fields.repository,
			}
			if err := s.ListLocation(tt.args.req, tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("service.ListLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_ReadLocation(t *testing.T) {
	type fields struct {
		repository repository.Repository
	}
	type args struct {
		ctx context.Context
		req *pb.ReadLocationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ReadLocationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repository: tt.fields.repository,
			}
			got, err := s.ReadLocation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ReadLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.ReadLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateLocation(t *testing.T) {
	type fields struct {
		repository repository.Repository
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateLocationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateLocationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repository: tt.fields.repository,
			}
			got, err := s.UpdateLocation(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UpdateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
