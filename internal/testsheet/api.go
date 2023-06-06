package album

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/sugaml/qa-api/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/albums/<id>", res.get)
	//	r.Get("/albums", res.query)

	r.Use(authHandler)

	// the following endpoints require a valid JWT
	// r.Post("/albums", res.create)
	// r.Put("/albums/<id>", res.update)
	// r.Delete("/albums/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	album, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(album)
}
