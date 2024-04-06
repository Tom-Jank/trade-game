package routes

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/middleware"
	"github.com/go-chi/chi/v5"
)

func CointossRoutes(handler *middleware.CointossHandler) http.Handler {
    r:= chi.NewRouter()
    r.Get("/{side}", handler.Toss)
    return r
}

