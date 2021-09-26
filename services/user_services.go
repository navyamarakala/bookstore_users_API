package services

import (
	"log"
	"time"

	"github.com/bookstore_users_API/utility/errors"

	"github.com/bookstore_users_API/model"
)

func CreateUser(user model.User) (*model.User, *errors.RestErrors) {
	if err := user.ValidateEmail(); err != nil {
		return nil, err
	}
	if saveerr := user.Save(); saveerr != nil {
		return nil, saveerr
	}
	return &user, nil
}

/*func GetUsers() (*model.User, *errors.RestErrors) {
	result, err := model.GetUsers()
	if err != nil {
		return nil, err
	}
	return result, nil
}*/

func GetUser(id int64) (*model.User, *errors.RestErrors) {
	var user model.User
	user.Id = id
	log.Println(id)
	err := user.GetUser()
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func EditUser(isPartial bool, user model.User) (*model.User, *errors.RestErrors) {
	current, geterr := GetUser(user.Id)
	if geterr != nil {
		return nil, geterr
	}
	if isPartial {
		if user.First_Name != "" {
			current.First_Name = user.First_Name
		}
		if user.Last_Name != "" {
			current.Last_Name = user.Last_Name
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.Email = user.Email
		current.First_Name = user.First_Name
		current.Last_Name = user.Last_Name

	}
	current.ModifiedDate = time.Now().UTC()
	err := current.EditedUser()
	if err != nil {
		return nil, err
	}

	return current, nil

}

func DeleteUser(id int64) (*model.User, *errors.RestErrors) {
	var user model.User
	user.Id = id

	err := user.DeleteUser()
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func GetUsersByField(lastName string) ([]model.User, *errors.RestErrors) {
	var user model.User
	user.Last_Name = lastName
	log.Printf("from services %s", user.Last_Name)
	data, err := user.GetDatabyField()
	if err != nil {
		return nil, err
	}
	return data, nil
}
