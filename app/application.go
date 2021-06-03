package app

import "github.com/gin-gonic/gin"

var (
	r = gin.Default()
)

func StartApplication() {
	mapUrls()
	r.Run(":8080")

}
