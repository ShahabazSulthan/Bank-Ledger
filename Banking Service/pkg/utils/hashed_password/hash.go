package hash

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type IHashPassword interface {
	HashedPassword(password string) string
	ComparePassword(hashedPassword string, plainPassword string) error
}

type HashUtil struct{}

func NewHashUtil() IHashPassword {
	return &HashUtil{}
}

func (h *HashUtil) HashedPassword(password string) string {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in Hashing Password: ", err)
	}

	fmt.Println(HashedPassword)
	return string(HashedPassword)
}

func (h *HashUtil) ComparePassword(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}
