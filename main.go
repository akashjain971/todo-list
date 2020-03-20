package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos/", handler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
