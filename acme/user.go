package acme

import (
	"crypto"
	"github.com/go-acme/lego/v4/registration"
)

type User struct {
	email        string
	key          crypto.PrivateKey
	registration *registration.Resource
}

func NewUser(email string, key crypto.PrivateKey) User {
	return User{
		email: email,
		key:   key,
	}
}

func NewRegisteredUser(email string, key crypto.PrivateKey, registration registration.Resource) User {
	return User{
		email:        email,
		key:          key,
		registration: &registration,
	}
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func (u *User) GetRegistration() *registration.Resource {
	return u.registration
}
