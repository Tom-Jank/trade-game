package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tom-Jank/trade-game/repository"
)

type UserHandler struct {
    userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
    return &UserHandler{userRepository: userRepository}
}

func (uh *UserHandler) HandleGetById(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)   
        w.Write([]byte("Bad parameters provided"))
        return
    }
    
    u, err := uh.userRepository.GetById(uint(id)) 

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)   
        w.Write([]byte("User with such ID doesn't exist"))
        return
    }
    jsonResponder(w, u)
}

// function to respond with simpl json in case I dont need to redner with templ
func jsonResponder(w http.ResponseWriter, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}

