package rest

import "net/http"

type Server struct {
	router *Router
}

func NewServer() *Server {
	return &Server{
		router: nil,
	}
}

func (this *Server) SetRouter(router *Router) {
	this.router = router
}

func (this *Server) Run() {
	handler := http.HandlerFunc(this.router.HandlerFunc())
	http.ListenAndServe(":8080", handler)
}
