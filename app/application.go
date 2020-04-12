package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication function is the starting point of whole application
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
