package main

import(
	"github.com/leak-test/mux"
	"net/http"
)

var router = mux.NewRouter()

func acceptsRouter(r *mux.Router) {
	panic("not implemented")
}

func main() {
    http.Handle("/", router)
}
