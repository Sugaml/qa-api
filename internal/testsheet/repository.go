package testsheet

import (
	"context"

	"github.com/sugaml/qa-api/internal/entity"
	"github.com/sugaml/qa-api/pkg/dbcontext"
	"github.com/sugaml/qa-api/pkg/log"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id string) (entity.TestSheet, error)
	// Count returns the number of albums.
	Count(ctx context.Context) (int, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.TestSheet, error)
	// Create saves a new album in the storage.
	Create(ctx context.Context, req entity.TestSheet) error
	// Update updates the album with given ID in the storage.
	Update(ctx context.Context, req entity.TestSheet) error
	// Delete removes the album with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the album with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.TestSheet, error) {
	var data entity.TestSheet

	return data, nil
}

// Create saves a new album record in the database.
// It returns the ID of the newly inserted album record.
func (r repository) Create(ctx context.Context, req entity.TestSheet) error {
	return nil
}

// Update saves the changes to an album in the database.
func (r repository) Update(ctx context.Context, req entity.TestSheet) error {
	return nil
}

// Delete deletes an album with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r repository) Count(ctx context.Context) (int, error) {
	return 0, nil
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.TestSheet, error) {
	return nil, nil
}
