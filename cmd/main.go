package main

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/db"
	"github.com/Tom-Jank/trade-game/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load(".env")
    db.Setup()
    r := chi.NewRouter()
    r.Mount("/", routes.RouterHello())
    http.ListenAndServe(":8080", r)
}
