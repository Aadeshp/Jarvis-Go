package rest

import (
	"net/http"
	"log"
)

type Router struct {
	Routes *Trie
}

func NewRouter(routes ...*Route) *Router {
	router := &Router{
		Routes: NewTrie(),
	}

	for _, route := range routes {
		router.Routes.Insert([]byte(route.Path), route)
	}

	return router
}

func (this *Router) AddRoute(route *Route) {
	this.Routes.Insert([]byte(route.Path), route)
}

func (this *Router) AddRoutes(prefix string, routes []*Route) {
	for _, route := range routes {
		route.Path = prefix + route.Path
		this.Routes.Insert([]byte(route.Path), route)
	}
}

func (this Router) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		request := NewRequest(req)
		writer := ResponseWriter{w}

		route := this.findRoute(request.Method, request.URL.RequestURI())
		if route == nil {
			return
		}
		log.Println("Listening...")
		handler := route.Handler
		handler(writer, request)
	}
}

func (this Router) findRoute(httpMethod string, path string) *Route {
	return this.Routes.Find(httpMethod, []byte(path))
}
