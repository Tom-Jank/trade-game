package routes

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/middleware"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(handler *middleware.UserHandler) http.Handler {
    r:= chi.NewRouter()
    r.Get("/{id}", handler.HandleGetById)
    return r
}
