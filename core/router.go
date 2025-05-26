package core

import (
	"fmt"
	"nroute/core/n_http"
	"strings"
)

type Router struct {
	routerTree *RouterTree
}

func NewRouter() *Router {
	return &Router{
		routerTree: NewRouterTree(),
	}
}

func (r *Router) addPathHandler(method n_http.Method, path string, h n_http.Handler) {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	r.routerTree.AddRoute(method, segments, h)
}

func (r *Router) GET(path string, h n_http.Handler) {
	r.addPathHandler(n_http.GET, path, h)
}

func (r *Router) find(path string) (*RouterTreeNode, error) {
	segments := strings.Split(strings.Trim(path, "/"), "/")

	node := r.routerTree.FindNode(segments)

	if node == nil {
		return nil, fmt.Errorf("path not found: %s", path)
	}

	return node, nil
}

func (r *Router) Use(method n_http.Method, path string) {
	node, err := r.find(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	methodHandlers := node.handlers[method]

	if len(methodHandlers) <= 0 {
		fmt.Println(fmt.Errorf("method '%s' not allowed for path '%s'", method, path))
		return
	}

	for _, h := range methodHandlers {
		h(path)
	}
}

// func (r *Router) UseRoute(path string) {
// 	if r.routerTree == nil {
// 		return
// 	}

// 	segments := strings.Split(strings.Trim(path, "/"), "/")
// 	var routerNode *RouterTree
// 	if len(segments) <= 1 && segments[0] == "" {
// 		routerNode = r.routerTree
// 	} else {
// 		routerNode = r.match(segments)
// 	}

// 	if routerNode == nil {
// 		return
// 	}

// 	for _, handler := range routerNode.handlers {
// 		handler(path)
// 	}
// }

// func (r *Router) match(segments []string) *RouterTree {
// 	if r.routerTree == nil {
// 		return nil
// 	}

// 	node := r.routerTree
// 	for _, seg := range segments {
// 		node = node.children[seg]
// 	}

// 	return node
// }
