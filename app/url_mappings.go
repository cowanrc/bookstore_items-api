package app

import (
	"bookstore_items-api/controllers/ping"
	"bookstore_items-api/controllers/users"
)

func mapUrls() {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:user_id", users.GetUser)
	// r.GET("/users/search", users.SearchUser)
	r.POST("/users", users.CreateUser)
}
