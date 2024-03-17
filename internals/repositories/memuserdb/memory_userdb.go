package memuserdb

import (
	"chapar/internals/core/domain"
	"log"
)

type MemoryUserDB struct {
	db []domain.User
}

func NewMemoryUserDB() *MemoryUserDB {
	return &MemoryUserDB{db: make([]domain.User, 0)}
}

func (d *MemoryUserDB) Create(newUser domain.User) (domain.User, error) {
	newUser.ID = uint(len(d.db) + 1)
	d.db = append(d.db, newUser)
	log.Println("created : ", d.db)
	return d.db[len(d.db)-1], nil
}
func (d *MemoryUserDB) GetUser(c domain.Credentials) (domain.User, error) {
	for i := range d.db {
		if d.db[i].UserName == c.UserName {
			return d.db[i], nil
		}
	}
	log.Println("searched : ", d.db)
	return domain.User{}, domain.UserNotFound{UserName: c.UserName}
}
