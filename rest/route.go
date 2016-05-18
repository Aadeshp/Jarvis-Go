package rest

//import "net/http"

type HandlerFunc func(ResponseWriter, *Request)

type Route struct {
	HttpMethod string
	Path       string
	Handler    HandlerFunc
}

func Get(path string, handler HandlerFunc) *Route {
	return &Route{
		HttpMethod: "GET",
		Path:       path,
		Handler:    handler,
	}
}

func Post(path string, handler HandlerFunc) *Route {
	return &Route{
		HttpMethod: "POST",
		Path:       path,
		Handler:    handler,
	}
}

func Put(path string, handler HandlerFunc) *Route {
	return &Route{
		HttpMethod: "PUT",
		Path:       path,
		Handler:    handler,
	}
}
