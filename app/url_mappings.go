package app

import "bookstore_items-api/controllers"

func mapUrls() {
	r.GET("/ping", controllers.Ping)

	r.GET("/users/:user_id", controllers.GetUser)
	// r.GET("/users/search", controllers.SearchUser)
	r.POST("/users", controllers.CreateUser)
}
