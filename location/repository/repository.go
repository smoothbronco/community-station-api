package repository

import (
	"context"
	"database/sql"

	"github.com/smoothbronco/community-station-api/location/pb"
	"gorm.io/gorm"
)

type Repository interface {
	InsertLocation(ctx context.Context, input *pb.LocationInput) (int64, error)
	SelectLocationByID(ctx context.Context, id int64) (*pb.Location, error)
	UpdateLocation(ctx context.Context, id int64, input *pb.LocationInput) error
	DeleteLocation(ctx context.Context, id int64) error
	SelectAllLocations() (*sql.Rows, error)
}

type mysqlRepo struct {
	db *gorm.DB
}

func NewMysqlRepo(db *gorm.DB) (Repository, error) {
	return &mysqlRepo{db}, nil
}

// DeleteLocation implements Repository
func (r *mysqlRepo) DeleteLocation(ctx context.Context, id int64) error {
	if err := r.db.Delete(&pb.Location{}, id).Error; err != nil {
		return err
	}
	return nil
}

// InsertLocation implements Repository
func (r *mysqlRepo) InsertLocation(ctx context.Context, input *pb.LocationInput) (int64, error) {
	location := pb.Location{Name: input.GetName(), Latitude: input.GetLatitude(), Longitude: input.GetLongitude(), Note: input.GetNote(), MapUrl: input.GetMapUrl()}
	if err := r.db.Create(&location).Error; err != nil {
		return 0, err
	}
	return location.GetId(), nil
}

// SelectAllLocations implements Repository
func (r *mysqlRepo) SelectAllLocations() (*sql.Rows, error) {
	locations, err := r.db.Model(&pb.Location{}).Rows()
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// SelectLocationByID implements Repository
func (r *mysqlRepo) SelectLocationByID(ctx context.Context, id int64) (*pb.Location, error) {
	var location *pb.Location
	if err := r.db.First(&location, id).Error; err != nil {
		return nil, err
	}
	return location, nil
}

// UpdateLocation implements Repository
func (r *mysqlRepo) UpdateLocation(ctx context.Context, id int64, input *pb.LocationInput) error {

	result := r.db.Model(pb.Location{}).Where("id = ?", id).Updates(
		pb.Location{
			Name:      input.GetName(),
			Latitude:  input.GetLatitude(),
			Longitude: input.GetLongitude(),
			Note:      input.GetNote(),
			MapUrl:    input.GetMapUrl(),
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
