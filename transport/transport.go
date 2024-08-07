package transport

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/udemy/go-angular/endpoint"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	return r
}
