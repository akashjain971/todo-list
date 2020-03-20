package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/todos/", handleGetRequest)
	router.GET("/todos/:task", handleGetRequest)

	router.POST("/todos/:task", handlePostRequest)

	router.PATCH("/todos/:task", handlePatchRequest)

	router.DELETE("/todos/", handleDeleteRequest)
	router.DELETE("/todos/:task", handleDeleteRequest)

	log.Fatal(http.ListenAndServe(":9999", router))
}
