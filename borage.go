package borage

import (
	"net/http"
)

// Borage struct
type Borage struct {
	router          *Router
	notFoundHandler http.HandlerFunc
	Debug           bool
}
type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(r.URL.Path))
}

// New func return Borage
func New() *Borage {
	b := &Borage{}
	b.router = NewRouter(b)
	b.Debug = true
	return b
}

func (b *Borage) Handle(method, path string, handle http.HandlerFunc) {
	if path[0] != '/' {
		panic("path must start with /")
	}
	b.router.tree.addNode(method, path, handle)
}

func (b *Borage) GET(path string, handle http.HandlerFunc) {
	b.Handle("GET", path, handle)
}

func (b *Borage) POST(path string, handle http.HandlerFunc) {
	b.Handle("POST", path, handle)
}

func (b *Borage) PUT(path string, handle http.HandlerFunc) {
	b.Handle("PUT", path, handle)
}

func (b *Borage) DELETE(path string, handle http.HandlerFunc) {
	b.Handle("DELETE", path, handle)
}

func (b *Borage) Start(addr string) {
	server := &Server{}
	http.ListenAndServe(addr, server)
}
