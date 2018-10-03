package borage

import (
	"net/http"
	"net/url"
	"strings"
)

// Node struct
type Node struct {
	component string
	children  []*Node
	methods   map[string]http.HandlerFunc
	isParam   bool
}

func (n *Node) addNode(method, path string, handler http.HandlerFunc) {
	componets := strings.Split(path, "/")[1:]
	component := componets[0]
	curNode := n
	aNode := n
	for {
		aNode = nil
		for _, child := range curNode.children {
			if child.component == component {
				aNode = child
				break
			}
		}
		if aNode == nil {
			aNode = &Node{
				component: component,
				children:  make([]*Node, 0),
				methods:   make(map[string]http.HandlerFunc),
				isParam:   false,
			}
			if component[0] == ':' {
				aNode.isParam = true
			}
			curNode.children = append(curNode.children, aNode)
		}
		curNode = aNode
		if len(componets) == 1 {
			curNode.methods[method] = handler
			break
		}
		componets = componets[1:]
		component = componets[0]
	}
}

func (n *Node) searchNode(path string, params url.Values) *Node {
	components := strings.Split(path, "/")[1:]
	component := components[0]
	curNode := n
	for {
		for _, child := range curNode.children {
			if child.component == component || child.isParam {
				if child.isParam && params != nil {
					params.Add(child.component[1:], component)
				}
				curNode = child
				break
			}
		}
		if len(components) == 1 {
			break
		}
		components = components[1:]
		component = components[0]
	}
	return curNode
}
