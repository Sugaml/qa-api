package healthcheck

import (
	"net/http"
	"testing"

	"github.com/sugaml/qa-api/internal/test"
	"github.com/sugaml/qa-api/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	RegisterHandlers(router, "0.9.0")
	test.Endpoint(t, router, test.APITestCase{
		Name: "ok", Method: "GET", URL: "/healthcheck", Body: "", Header: nil, WantStatus: http.StatusOK, WantResponse: `"OK 0.9.0"`,
	})
}
