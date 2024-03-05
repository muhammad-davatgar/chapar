package apperrors

import "fmt"

var (
	ListenerNotFound = fmt.Errorf("user offline")
	UserNotFound     = fmt.Errorf("user not found")
)
