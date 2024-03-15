package middleware

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/views"
	"github.com/a-h/templ"
)

func HandleGreeting(w http.ResponseWriter, r *http.Request) {
    templ.Handler(views.Layout()).ServeHTTP(w, r)
}

func HandleEntry(w http.ResponseWriter, r *http.Request) {
    templ.Handler(views.MainPage()).ServeHTTP(w, r)
}
