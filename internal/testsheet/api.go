package testsheet

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/sugaml/qa-api/internal/errors"
	"github.com/sugaml/qa-api/pkg/log"
	"github.com/sugaml/qa-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/testsheet/<id>", res.get)
	r.Get("/testsheets", res.query)

	r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/testsheets", res.create)
	r.Put("/testsheets/<id>", res.update)
	r.Delete("/testsheets/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	testsheet, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(testsheet)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	albums, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = albums
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateTestSheetRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	testsheet, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(testsheet, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateTestSheetRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	testsheet, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(testsheet)
}

func (r resource) delete(c *routing.Context) error {
	testsheet, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(testsheet)
}
