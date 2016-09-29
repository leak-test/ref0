package main

import(
	"github.com/leak-test/mux"
	"net/http"
)

func Serve() {
    http.Handle("/", router)
}
