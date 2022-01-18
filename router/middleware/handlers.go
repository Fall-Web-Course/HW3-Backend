package middleware

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Fall-Web-Course/HW3/utils"
	"github.com/gin-gonic/gin"
)

var requests map[string]int = make(map[string]int)
var RATE_LIMIT int = getRateLimit() 

func RateLimit(c *gin.Context) {
	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	if requests[ip] == RATE_LIMIT {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H {
			"Message": "To many requests",
		})
		return
	}
	requests[ip] += 1
	go requestTimeWatcher(ip)

	c.Next()
}

func requestTimeWatcher(ip string) {
	time.Sleep(60 * time.Second)
	requests[ip] -= 1
}

func getRateLimit() int {
	value, _ := strconv.Atoi(utils.Getenv("RATE_LIMIT", "10"))
	return value
}