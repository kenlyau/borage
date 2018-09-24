package borage

import (
	"net/http"
	"strings"
)

// Node struct
type Node struct {
	component string
	children  map[string]*Node
	methods   map[string]http.HandlerFunc
	isParamed bool
}

func (n *Node) addNode(method, path string, handler http.HandlerFunc) {
	componets := strings.Split(path, "/")[1:]
	component := componets[0]
	curNode := n
	for len(componets) > 0 {
		if curNode.children[component] != nil {
			curNode = curNode.children[component]
		} else {
			newNode := &Node{
				component: component,
				children:  make(map[string]*Node),
				methods:   make(map[string]http.HandlerFunc),
				isParamed: false,
			}
			curNode.children[component] = newNode
			curNode = curNode.children[component]
		}
		if len(componets) == 1 {
			curNode.methods[method] = handler
		}
		componets = componets[1:]
		component = componets[0]
	}
}

func (n *Node) searchNode(method, path string) *Node {
	components := strings.Split(path, "/")[1:]
	component := components[0]
	curNode := n
	for len(components) > 0 {
		if curNode.children[component] != nil {
			curNode = curNode.children[component]
		}
		if len(components) == 1 {
			break
		}
		components = components[1:]
		component = components[0]
	}
	if curNode.component == component {
		return curNode
	}
	return nil
}
