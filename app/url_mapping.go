package app

import "github.com/bookstore_users_API/controllers"

func mapUrls() {

	router.POST("/user", controllers.CreateUser)
	//router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUser)
	router.PUT("/user/:id", controllers.EditUser)
	router.PATCH("/user/:id", controllers.EditUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
	router.GET("/user/search", controllers.GetUsersByField)

}
