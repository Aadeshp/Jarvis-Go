package route

import (
	"Jarvis-Go/rest"
	"log"
	"io/ioutil"
)

var WebhookRoutes = []*rest.Route{
	rest.Get("/", func(writer rest.ResponseWriter, req *rest.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("Error!")
		}
		log.Println(string(body))
	}),

	rest.Post("/", func(writer rest.ResponseWriter, req *rest.Request) {
		// Handle webhook POST request
	}),
}
