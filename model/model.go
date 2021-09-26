package model

import (
	"strings"
	"time"

	"github.com/bookstore_users_API/utility/errors"
)

type User struct {
	Id           int64     `json:"id"`
	First_Name   string    `json:"firstName"`
	Last_Name    string    `json:"lastName"`
	Email        string    `json:"email"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

type Users []struct {
	User
}

func (user User) ValidateEmail() *errors.RestErrors {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid mail Id")
	}

	return nil
}
