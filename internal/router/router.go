package router

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Route struct {
	handler     http.Handler
	middlewares []Middleware
}

type Router struct {
	mux    *http.ServeMux
	routes map[string]*Route
}

func NewRouter() *Router {
	return &Router{
		mux:    http.NewServeMux(),
		routes: make(map[string]*Route),
	}
}

func (cr *Router) Handle(path string, handler http.Handler, middlewares ...Middleware) {
	cr.routes[path] = &Route{
		handler:     handler,
		middlewares: middlewares,
	}
	cr.mux.Handle(path, cr.applyMiddlewares(path))
}

func (cr *Router) applyMiddlewares(path string) http.Handler {
	route := cr.routes[path]
	return ChainMiddlewares(route.handler, route.middlewares...)
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
