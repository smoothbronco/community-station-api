package repository

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/smoothbronco/community-station-api/location/pb"
	"gorm.io/gorm"
)

func TestNewMysqlRepo(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMysqlRepo(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMysqlRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMysqlRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRepo_DeleteLocation(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  int64
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
			r := &mysqlRepo{
				db: tt.fields.db,
			}
			if err := r.DeleteLocation(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepo.DeleteLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlRepo_InsertLocation(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx   context.Context
		input *pb.LocationInput
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
			r := &mysqlRepo{
				db: tt.fields.db,
			}
			got, err := r.InsertLocation(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepo.InsertLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mysqlRepo.InsertLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRepo_SelectAllLocations(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepo{
				db: tt.fields.db,
			}
			got, err := r.SelectAllLocations()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepo.SelectAllLocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRepo.SelectAllLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRepo_SelectLocationByID(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Location
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepo{
				db: tt.fields.db,
			}
			got, err := r.SelectLocationByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepo.SelectLocationByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlRepo.SelectLocationByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mysqlRepo_UpdateLocation(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mysqlRepo{
				db: tt.fields.db,
			}
			if err := r.UpdateLocation(tt.args.ctx, tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("mysqlRepo.UpdateLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
