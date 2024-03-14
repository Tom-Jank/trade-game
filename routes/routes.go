package routes

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/middleware"
	"github.com/go-chi/chi/v5"
)

func RouterHello() http.Handler {
    r := chi.NewRouter()
    r.Get("/", middleware.HandleEntry)
    r.Get("/{id}", middleware.HandleGreet)
    return r
}
