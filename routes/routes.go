package routes

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/middleware"
	"github.com/go-chi/chi/v5"
)

func RouterHello() http.Handler {
    r := chi.NewRouter()
    r.Get("/", middleware.HandleGreeting)
    r.Get("/main-page", middleware.HandleEntry)
    return r
}
