package middleware

import (
	"net/http"

	"github.com/Tom-Jank/trade-game/repository"
)

type CointossHandler struct {
    cointossRepository repository.CointossRepository
}

func NewCointossHandler(
    cointossRepository repository.CointossRepository) *CointossHandler {

    return &CointossHandler{cointossRepository: cointossRepository}
}

func (ch *CointossHandler) Toss(w http.ResponseWriter, r *http.Request) {
    // side = sideConv(r.PathValue("side"))
    result := ch.cointossRepository.TossACoin()
    jsonResponder(w, result)
}

func sideConv(side string) bool {
    if side == "HEAD" {
        return true
    } else {
        return false
    }
}
