package main

import(
	"github.com/leak-test/mux"
	"net/http"
)

var router = mux.NewRouter()

func Serve() {
    http.Handle("/", router)
}
