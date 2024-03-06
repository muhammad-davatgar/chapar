package domain

type User struct {
	ID      uint
	Reciver chan Message
}
