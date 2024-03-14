package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	repo "github.com/Tom-Jank/trade-game/repository"
	"github.com/Tom-Jank/trade-game/views"
	"github.com/a-h/templ"
)

func HandleGreet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)   
        w.Write([]byte("Bad parameters provided"))
        return
    }
    u, err := repo.FindUser(uint(id)) 
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)   
        w.Write([]byte("User with such ID doesn't exist"))
        return
    }
    jsonResponder(w, u)
    //templ.Handler(views.MainPage(u.Name)).ServeHTTP(w, r)
}

func HandleEntry(w http.ResponseWriter, r *http.Request) {
    templ.Handler(views.Layout()).ServeHTTP(w, r)
}

// function to respond with simpl json in case I dont need to redner with templ
func jsonResponder(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}
