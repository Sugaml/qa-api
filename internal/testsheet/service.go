package testsheet

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sugaml/qa-api/internal/entity"
	"github.com/sugaml/qa-api/pkg/log"
)

// Service encapsulates usecase logic for testsheets.
type Service interface {
	Get(ctx context.Context, id string) (TestSheet, error)
	Query(ctx context.Context, offset, limit int) ([]TestSheet, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateTestSheetRequest) (TestSheet, error)
	Update(ctx context.Context, id string, input UpdateTestSheetRequest) (TestSheet, error)
	Delete(ctx context.Context, id string) (TestSheet, error)
}

// TestSheet represents the data about an testsheet.
type TestSheet struct {
	entity.TestSheet
}

// CreateTestSheetRequest represents an testsheet creation request.
type CreateTestSheetRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateTestsheetRequest fields.
func (m CreateTestSheetRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateTestsheetRequest represents an testsheet update request.
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

// NewService creates a new testsheet service.
func NewService(repo Repository, logger log.Logger) Service {
	return nil
}

// Get returns the album with the specified the testsheet ID.
func (s service) Get(ctx context.Context, id string) (TestSheet, error) {
	album, err := s.repo.Get(ctx, id)
	if err != nil {
		return TestSheet{}, err
	}
	return TestSheet{album}, nil
}

// Create creates a new testsheet.
func (s service) Create(ctx context.Context, req CreateTestSheetRequest) (TestSheet, error) {
	if err := req.Validate(); err != nil {
		return TestSheet{}, err
	}
	id := string(0)
	now := time.Now()
	err := s.repo.Create(ctx, entity.TestSheet{
		ID:        id,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return TestSheet{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the testheet with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateTestSheetRequest) (TestSheet, error) {
	if err := req.Validate(); err != nil {
		return TestSheet{}, err
	}

	testheet, err := s.Get(ctx, id)
	if err != nil {
		return testheet, err
	}
	testheet.Name = req.Name
	testheet.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, testheet.TestSheet); err != nil {
		return testheet, err
	}
	return testheet, nil
}

// Delete deletes the testheet with the specified ID.
func (s service) Delete(ctx context.Context, id string) (TestSheet, error) {
	testheet, err := s.Get(ctx, id)
	if err != nil {
		return TestSheet{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return TestSheet{}, err
	}
	return testheet, nil
}

// Count returns the number of albums.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the testsheets with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]TestSheet, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []TestSheet{}
	for _, item := range items {
		result = append(result, TestSheet{item})
	}
	return result, nil
}
