package repository

import (
	"math/rand"
	"time"
)

type CointossRepository interface {
    TossACoin() bool
}

type CointossRepositoryImpl struct {
}

func NewCointossRepositoryImpl() *CointossRepositoryImpl {
    return &CointossRepositoryImpl{}
}

func (cr *CointossRepositoryImpl) TossACoin() bool {
    result := randbool()
    // add result to a db with user name
    // return result
    return result
}

func randbool() bool {
    seed := time.Now().UnixNano()
    rng := rand.New(rand.NewSource(seed))
    return rng.Uint64()&(1<<63) != 0
}

