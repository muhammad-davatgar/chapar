package domain_test

import (
	"chapar/internals/core/domain"
	"testing"
)

func TestPasswordEncryption(t *testing.T) {
	pass := "password123"
	cred := domain.Credentials{UserName: "mmddvg", PassWord: pass}

	user, err := cred.EncryptPassword()
	if err != nil {
		t.Errorf("error while encrypting : %v ", err.Error())
	}

	res, err := user.VerifyPassword(pass)
	if err != nil {
		t.Errorf("err while decrypting : %v", err.Error())
	}

	if !res {
		t.Errorf("got false while comparing ; encrypted password : %v", user.PassWord)
	}
}
