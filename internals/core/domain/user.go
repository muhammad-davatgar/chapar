package domain

import (
	"chapar/shared"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"

	"github.com/golang-jwt/jwt"
)

type HubUser struct {
	ID      uint
	Reciver chan Message
}

type EntryCredentials struct {
	UserName string `validate:"required,min=3" json:"user_name"`
	PassWord string `validate:"required,min=8" json:"password"`
}

func (v *Validator) Credentials(d EntryCredentials) (Credentials, error) {
	credentials := Credentials{}

	err := v.validate.Struct(d)
	if err != nil {
		return credentials, err
	}

	credentials.UserName = d.UserName
	credentials.PassWord = d.PassWord

	return credentials, nil

}

type Credentials struct {
	UserName string
	PassWord string
}

func generateMd5(inp string) []byte {
	h := md5.New()
	h.Write([]byte(inp))
	return h.Sum(nil)
}

func (c *Credentials) EncryptPassword() (User, error) {
	// using md5 to get a fixed size password

	fixedPassword := generateMd5(c.PassWord)
	block, err := aes.NewCipher(fixedPassword)
	if err != nil {
		return User{}, err
	}
	plainText := []byte(c.UserName)
	cfb := cipher.NewCFBEncrypter(block, fixedPassword)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	p := base64.StdEncoding.EncodeToString(cipherText)
	return User{UserName: c.UserName, PassWord: p}, nil
}

type User struct {
	ID       uint
	UserName string
	PassWord string
}

func (u *User) RefreshToken() string {
	claims := shared.NewClaims(u.ID, u.UserName, shared.RefreshToken)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte("secret")) // TODO : env var

	return t
}

func (u *User) VerifyPassword(entryPassword string) (bool, error) {
	fixedPassword := generateMd5(entryPassword)
	block, err := aes.NewCipher(fixedPassword)
	if err != nil {
		return false, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(u.PassWord)
	if err != nil {
		return false, err
	}

	cfb := cipher.NewCFBDecrypter(block, fixedPassword)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText) == u.UserName, nil
}
