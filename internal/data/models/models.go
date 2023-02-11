package models

import (
	"database/sql"
	"errors"

	"github.com/SergioRosello/greenlight/internal/data"
)

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Movies interface {
		Insert(movie *data.Movie) error
		Get(id int64) (*data.Movie, error)
		GetAll(title string, genres []string, filters data.Filters) ([]*data.Movie, data.Metadata, error)
		Update(movie *data.Movie) error
		Delete(id int64) error
	}
	Users interface {
		Insert(*data.User) error
		GetByEmail(email string) (*data.User, error)
		Update(movie *data.User) error
	}
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
	}
}

// Create a helper function which returns a Models instance containing the mock models
// only.
func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
		Users:  MockUserModel{},
	}
}
