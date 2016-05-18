package main

import (
	"Jarvis-Go/rest"
	"Jarvis-Go/route"
)

func main() {

	server := rest.NewServer()

	router := rest.NewRouter()
	router.AddRoutes("/webhook", route.WebhookRoutes)

	server.SetRouter(router)
	server.Run()
}
