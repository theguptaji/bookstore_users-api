package app

import (
	"github.com/gin-gonic/gin"
	"github.com/theguptaji/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

// StartApplication function is the starting point of whole application
func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8081")
}
