package borage

import "net/http"

// Router struct
type Router struct {
	tree   *Node
	borage *Borage
}

func NewRouter(b *Borage) *Router {
	node := &Node{
		component: "",
		children:  make(map[string]*Node),
		methods:   make(map[string]http.HandlerFunc),
	}
	return &Router{
		tree:   node,
		borage: b,
	}
}
