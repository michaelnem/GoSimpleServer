package core

import "nroute/core/n_http"

type Route struct {
	Path    string
	Handler n_http.Handler
}
