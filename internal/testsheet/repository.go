package testsheet

import (
	"context"

	"github.com/sugaml/qa-api/internal/entity"
	"github.com/sugaml/qa-api/pkg/dbcontext"
	"github.com/sugaml/qa-api/pkg/log"
)

// Repository encapsulates the logic to access testsheets from the data source.
type Repository interface {
	// Get returns the testsheet with the specified testsheet ID.
	Get(ctx context.Context, id string) (entity.TestSheet, error)
	// Count returns the number of testsheets.
	Count(ctx context.Context) (int, error)
	// Query returns the list of testsheets with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.TestSheet, error)
	// Create saves a new testsheet in the storage.
	Create(ctx context.Context, req entity.TestSheet) error
	// Update updates the testsheet with given ID in the storage.
	Update(ctx context.Context, req entity.TestSheet) error
	// Delete removes the testsheet with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new testsheet repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the testsheet with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.TestSheet, error) {
	var data entity.TestSheet

	return data, nil
}

// Create saves a new testsheet record in the database.
// It returns the ID of the newly inserted testsheet record.
func (r repository) Create(ctx context.Context, req entity.TestSheet) error {
	return nil
}

// Update saves the changes to an testsheet in the database.
func (r repository) Update(ctx context.Context, req entity.TestSheet) error {
	return nil
}

// Delete deletes an testsheet with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r repository) Count(ctx context.Context) (int, error) {
	return 0, nil
}

func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.TestSheet, error) {
	return nil, nil
}
