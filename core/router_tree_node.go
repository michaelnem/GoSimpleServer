package core

import "nroute/core/n_http"

type RouterTreeNode struct {
	pathSegement string
	children     map[string]*RouterTreeNode
	handlers     map[n_http.Method][]n_http.Handler
}

func NewRouterTreeNode(segment string) *RouterTreeNode {
	return &RouterTreeNode{
		children: make(map[string]*RouterTreeNode),
		handlers: make(map[n_http.Method][]n_http.Handler),
	}
}

func (node *RouterTreeNode) addHandler(method n_http.Method, h n_http.Handler) {
	if h == nil {
		return
	}
	node.handlers[method] = append(node.handlers[method], h)
}
