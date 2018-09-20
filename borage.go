package borage

type Borage struct {
	router          *Router
	notFoundHandler HandlerFunc
	Debug           bool
}
