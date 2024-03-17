package domain

import "fmt"

type UserNotFound struct {
	UserName string
}

func (e UserNotFound) Error() string {
	return fmt.Sprintf("user %v not found", e.UserName)
}

var (
	ErrListenerNotFound = fmt.Errorf("user offline")
)
