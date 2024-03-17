package auth

import (
	"chapar/internals/core/domain"
	"chapar/internals/core/ports"
	"fmt"
)

type AuthenticationService struct {
	userDB    ports.UserDB
	validator *domain.Validator
}

func NewAuthenticationService(userDB ports.UserDB) AuthenticationService {
	return AuthenticationService{userDB: userDB, validator: domain.NewValidator()}
}

func (s *AuthenticationService) SignUp(entryCredentials domain.EntryCredentials) (string, error) {
	credentials, err := s.validator.Credentials(entryCredentials)
	if err != nil {
		return "", err
	}

	user, err := credentials.EncryptPassword()
	if err != nil {
		return "", err
	}

	user, err = s.userDB.Create(user)
	if err != nil {
		return "", err
	}

	return user.RefreshToken(), nil

}

func (s *AuthenticationService) Login(entryCredentials domain.EntryCredentials) (string, error) {
	credentials, err := s.validator.Credentials(entryCredentials)
	if err != nil {
		return "", err
	}

	user, err := s.userDB.GetUser(credentials)
	if err != nil {
		return "", err
	}

	correct, err := user.VerifyPassword(credentials.PassWord)
	if err != nil {
		return "", err
	}

	if correct {
		return user.RefreshToken(), nil
	} else {
		return "", fmt.Errorf("incorrect password")
	}
}
