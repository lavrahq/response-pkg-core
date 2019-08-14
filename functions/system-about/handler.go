package function

import (
	"encoding/json"
	"net/http"
	"os"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Service is an instance of a running service.
type Service struct {
	Name    string
	Version string
	Data    []byte
}

// About stores information about
type About struct {
	Hostname string
	Services []Service
}

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	if req.Method != "GET" {
		return handler.Response{
			Body:       nil,
			StatusCode: http.StatusNotImplemented,
		}, nil
	}

	about := About{}

	if hostname, err := os.Hostname(); err != nil {
		about.Hostname = hostname
	}

	js, err := json.Marshal(about)

	return handler.Response{
		Body:       js,
		StatusCode: http.StatusOK,
	}, err
}
