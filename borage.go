package borage

import (
	"net/http"
)

// Borage struct
type Borage struct {
	router          *Router
	server          *Server
	notFoundHandler http.HandlerFunc
	Debug           bool
}
type Server struct {
	borage *Borage
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	tree := s.borage.router.tree
	node := tree.searchNode(r.Method, path)
	if node != nil {
		node.methods[r.Method](w, r)
	}
}

// New func return Borage
func New() *Borage {
	b := &Borage{}
	b.router = NewRouter(b)
	b.Debug = true
	server := &Server{borage: b}
	b.server = server
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
	http.ListenAndServe(addr, b.server)
}
