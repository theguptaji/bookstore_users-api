package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Ping return ping
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}