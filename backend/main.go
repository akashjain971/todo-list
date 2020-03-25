package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/todos/", handleGetRequest)
	router.GET("/todos/:uuid", handleGetRequestEx)

	router.POST("/todos/", handlePostRequest)

	router.PATCH("/todos/:uuid", handlePatchRequest)

	router.DELETE("/todos/", handleDeleteRequest)
	router.DELETE("/todos/:uuid", handleDeleteRequestEx)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
	}).Handler(router)
	log.Fatal(http.ListenAndServe(":9999", handler))
}
