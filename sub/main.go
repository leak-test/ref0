package main

import(
	"github.com/leak-test/mux"
	"net/http"
)

var router = mux.NewRouter()

func main() {
    http.Handle("/", router)
}
