package borage

import (
	"net/http"
)

type Borage struct {
	router          *Router
	notFoundHandler http.HandlerFunc
	Debug           bool
}
