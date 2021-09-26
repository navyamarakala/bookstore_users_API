package model

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bookstore_users_API/utility/errors"

	"github.com/bookstore_users_API/datasource/mysql/user_db"
)

var DbMap = make(map[int64]*User)
var (
	noUser            = "no rows in result set"
	insertStmt        = "insert into users(id,firstName,lastName,email,createdDate,modifiedDate) values(?,?,?,?,?,?);"
	getuserById       = "select * from users where id=?;"
	updateUser        = "update users set firstName=?,lastName=?,email=? where id=?;"
	deleteUser        = "delete from users where id=?"
	getDatabyLastName = "select * from users where lastName=?;"
)

func (user *User) Save() *errors.RestErrors {

	stmt, err := user_db.Db.Prepare(insertStmt)
	if err != nil {
		return errors.NewBaInternalServerErrordRequestError(err.Error())
	}
	defer stmt.Close()
	user.CreatedDate = time.Now().UTC()
	user.ModifiedDate = time.Now().UTC()
	_, exeerr := stmt.Exec(user.Id, user.First_Name, user.Last_Name, user.Email, user.CreatedDate, user.ModifiedDate)
	if exeerr != nil {
		if strings.Contains(exeerr.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError("Duplicate EmailId")
		}
		return errors.NewBaInternalServerErrordRequestError(fmt.Sprintf("unable to save the user!! Error is %s", exeerr.Error()))
	}
	return nil

}

func (user *User) GetUser() *errors.RestErrors {

	stmt, err := user_db.Db.Prepare(getuserById)
	if err != nil {
		return errors.NewBaInternalServerErrordRequestError(err.Error())
	}
	defer stmt.Close()
	data := stmt.QueryRow(user.Id)
	exeerr := data.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.CreatedDate, &user.ModifiedDate)
	if exeerr != nil {
		if strings.Contains(exeerr.Error(), noUser) {
			return errors.NewBadRequestError(fmt.Sprintf("user %d does not exist in DB", user.Id))
		}
		return errors.NewBaInternalServerErrordRequestError(fmt.Sprintf("unable to get the user!! Error is %s", exeerr.Error()))
	}
	return nil

}

func (user *User) EditedUser() *errors.RestErrors {
	stmt, err := user_db.Db.Prepare(updateUser)
	if err != nil {
		return errors.NewBaInternalServerErrordRequestError(err.Error())
	}
	defer stmt.Close()

	_, execerr := stmt.Exec(user.First_Name, user.Last_Name, user.Email, user.Id)
	if execerr != nil {
		if strings.Contains(execerr.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError("There is another record with the same email Id")
		}
		return errors.NewBaInternalServerErrordRequestError(fmt.Sprintf("unable to edit the user!! Error is %s", execerr.Error()))
	}
	return nil

}

func (user *User) DeleteUser() *errors.RestErrors {

	stmt, err := user_db.Db.Prepare(deleteUser)
	if err != nil {
		return errors.NewBaInternalServerErrordRequestError(err.Error())
	}
	defer stmt.Close()

	data, execerr := stmt.Exec(user.Id)
	if execerr != nil {
		return errors.NewBaInternalServerErrordRequestError(fmt.Sprintf("unable to delete the user!! Error is %s", execerr.Error()))
	}
	if data == nil {
		return errors.NewBaInternalServerErrordRequestError(fmt.Sprintf("No user with the id %d.. error is  %s", user.Id, execerr.Error()))
	}
	return nil

}

func (user *User) GetDatabyField() ([]User, *errors.RestErrors) {
	log.Printf(user.Last_Name)
	stmt, err := user_db.Db.Prepare(getDatabyLastName)
	if err != nil {
		return nil, errors.NewBaInternalServerErrordRequestError(err.Error())
	}
	defer stmt.Close()
	rows, queryerr := stmt.Query(user.Last_Name)
	if queryerr != nil {
		return nil, errors.NewBaInternalServerErrordRequestError(err.Error())
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		if err := rows.Scan(user.Id, user.First_Name, user.Last_Name, user.Email, user.CreatedDate); err != nil {
			return nil, errors.NewBaInternalServerErrordRequestError(err.Error())
		}
		results = append(results, *user)
	}
	return results, nil
}
