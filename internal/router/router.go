package router

import (
	"fmt"
	"net/http"
)

type Middleware = func(http.Handler) http.Handler

type Route struct {
	handler     http.Handler
	middlewares []Middleware
}

type Router struct {
	mux         *http.ServeMux
	routes      map[string]*Route
	middlewares []Middleware
}

type RequestMethod = int

const (
	GET RequestMethod = iota
	POST
	PATCH
	PUT
	DELETE
)

func NewRouter() *Router {
	return &Router{
		mux:         http.NewServeMux(),
		routes:      make(map[string]*Route),
		middlewares: []Middleware{},
	}
}

func (cr *Router) Handle(method RequestMethod, path string, handler http.Handler, middlewares ...Middleware) {
	var m string

	switch method {
	case 0:
		m = "GET"
	case 1:
		m = "POST"
	case 2:
		m = "PATCH"
	case 3:
		m = "PUT"
	case 4:
		m = "DELETE"
	default:
		m = "HEAD"
	}

	cr.routes[path] = &Route{
		handler:     handler,
		middlewares: middlewares,
	}

	request := fmt.Sprintf("%s %s", m, path)

	cr.mux.Handle(request, cr.applyMiddlewares(path))
}

func (cr *Router) RegisterMiddlewares(middlewares ...Middleware) {

	cr.middlewares = middlewares

}

func (cr *Router) applyMiddlewares(path string) http.Handler {
	route := cr.routes[path]
	middlewares := append(cr.middlewares, route.middlewares...)

	return ChainMiddlewares(route.handler, middlewares...)
}

func (cr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cr.mux.ServeHTTP(w, r)
}

func ChainMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
