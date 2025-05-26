package core

import (
	"nroute/core/n_http"
)

type RouterTree struct {
	root *RouterTreeNode
}

func NewRouterTree() *RouterTree {
	return &RouterTree{
		root: NewRouterTreeNode(""),
	}
}

func (rt *RouterTree) FindNode(segments []string) *RouterTreeNode {
	node := rt.root

	for _, segment := range segments {
		child, ok := node.children[segment]
		if !ok {
			return nil
		}
		node = child
	}

	return node
}

func (rt *RouterTree) AddRoute(method n_http.Method, segments []string, h n_http.Handler) {
	if len(segments) <= 1 && segments[0] == "" {
		rt.root.addHandler(method, h)
		return
	}

	node := rt.addPathSegments(segments)
	node.addHandler(method, h)
}

func (rt *RouterTree) addPathSegments(segments []string) *RouterTreeNode {
	node := rt.root

	for _, segment := range segments {
		child, ok := node.children[segment]
		if !ok {
			child = NewRouterTreeNode(segment)
			node.children[segment] = child

			child.pathSegement = segment
		}
		node = child
	}

	return node
}
