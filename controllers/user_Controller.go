package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bookstore_users_API/model"
	"github.com/bookstore_users_API/services"
	"github.com/bookstore_users_API/utility/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user model.User
	log.Println("user is", user)

	//this line is the replacement for reading the reqest and unmarshaling the req to struct
	if err := c.ShouldBindJSON(&user); err != nil {
		errors.NewBadRequestError("Bad JSON Request")
	}

	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		log.Fatal(err)
	}*/
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Println("user after unmarshal", user)

	c.JSON(http.StatusCreated, result)

}

/*func GetUsers(c *gin.Context) {
	results, err := services.GetUsers()
	if err != nil {
		c.String(http.StatusNotFound, err)
	}

	c.String(http.StatusAccepted, results)
}*/

func GetUser(c *gin.Context) {

	id1, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Please pass Interger"))
		return
	}
	log.Println("id is ", id1)
	userDetails, serviceerr := services.GetUser(id1)
	if serviceerr != nil {
		c.JSON(http.StatusBadRequest, serviceerr)
		return
	}

	c.JSON(http.StatusAccepted, userDetails)

}

func EditUser(c *gin.Context) {
	var user model.User
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Please pass Interger Value"))
		return
	}
	if userId <= 0 {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Please give valid integer"))
		return
	}
	if jsonerr := c.ShouldBindJSON(&user); jsonerr != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid JSON"))
	}
	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userId
	userDetails, resterr := services.EditUser(isPartial, user)
	if resterr != nil {
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	c.JSON(http.StatusAccepted, userDetails)
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid Json"))
		return
	}
	if userId <= 0 {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Please enter valid Interger ID"))
		return
	}

	_, resterr := services.DeleteUser(userId)
	if resterr != nil {
		c.JSON(http.StatusNotFound, resterr)
		return
	}

	c.JSON(http.StatusAccepted, map[string]string{"status": "deleted the user"})

}

func GetUsersByField(c *gin.Context) {
	lastName := c.Query("lastName")
	log.Println(lastName)
	data, err := services.GetUsersByField(lastName)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusAccepted, data)

}
