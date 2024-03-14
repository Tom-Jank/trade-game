package repo

import (
	"github.com/Tom-Jank/trade-game/db"
)

type User struct {
    ID uint `gorm:"primaryKey"`
    Name string
}

func FindUser(id uint) (User, error) {
    var user User
    if err := db.DBCon.First(&user, id).Error; err != nil {
        return user, err
    }
    return user, nil
}
