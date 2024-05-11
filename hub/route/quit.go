package route

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func quitRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", quitCore)
	return r
}

func quitCore(w http.ResponseWriter, r *http.Request) {
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(os.Interrupt)
}