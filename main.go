package main

import(
	"github.com/leak-test/mux"
	"net/http"
)

var router = mux.NewRouter()

func acceptsRouter(r *mux.Router) {
	panic("not implemented")
}

var myError = mux.SkipRouter

func main() {
    http.Handle("/", router)
}
