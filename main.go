package main

import (
	"log"
	"net/http"

	"go-api-test/infra/api"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", api.Init()))
}
