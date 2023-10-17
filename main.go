package main

import (
	"fmt"
	"main/controller"

	"net/http"

	"github.com/rs/cors"
)

func main() {
	fmt.Println("Start")
	mux := controller.Register()
	//go get github.com/rs/cors
	c := cors.New(cors.Options{ //setting up a CORS middleware for HTTP server
		AllowedOrigins: []string{"*"}, //allowedorigins option specifies which websuite are allowed to make requests to your server
		//*allows requests from any origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"}, //content-type in the request to tell about content
	})
	handler := c.Handler(mux) //combining the CORS handler with mux
	http.ListenAndServe("localhost:3000", handler)
	//we are listening for client request on port 3000
	fmt.Println("End")
}
