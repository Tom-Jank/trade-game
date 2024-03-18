package main

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/db"
	"github.com/Tom-Jank/trade-game/middleware"
	"github.com/Tom-Jank/trade-game/repository"
	"github.com/Tom-Jank/trade-game/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
    godotenv.Load(".env")
    db := db.Setup()
    r := chi.NewRouter()
    injectDependencies(db, r)
    http.ListenAndServe(":8080", r)
}

func injectDependencies(db *gorm.DB, r *chi.Mux) {
    userRepository := repository.NewUserRepositoryImpl(db)
    userHandler := middleware.NewUserHandler(userRepository)
    r.Mount("/api/user", routes.UserRoutes(userHandler))
    
    viewHandler := middleware.NewViewHandler()
    r.Mount("/", routes.ViewRoutes(viewHandler))
}
