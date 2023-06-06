package album

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/sugaml/qa-api/internal/entity"
)

// Service encapsulates usecase logic for albums.
type Service interface {
	Get(ctx context.Context, id string) (TestSheet, error)
	Query(ctx context.Context, offset, limit int) ([]TestSheet, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTestSheetRequest) (TestSheet, error)
	Update(ctx context.Context, id string, input UpdateTestSheetRequest) (TestSheet, error)
	Delete(ctx context.Context, id string) (TestSheet, error)
}

// Album represents the data about an album.
type TestSheet struct {
	entity.TestSheet
}

// CreateAlbumRequest represents an album creation request.
type CreateTestSheetRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateTestSheetRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateAlbumRequest represents an album update request.
type UpdateTestSheetRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateAlbumRequest fields.
func (m UpdateTestSheetRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new album service.
func NewService(repo Repository, logger log.Logger) Service {
	return nil
}
