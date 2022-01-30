package middleware

import (
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/Fall-Web-Course/HW3/authorization"
	"github.com/Fall-Web-Course/HW3/utils"
	"github.com/gin-gonic/gin"
)

var requests map[string]int = make(map[string]int)
var RATE_LIMIT int = getRateLimit()

func RateLimit(c *gin.Context) {
	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	if requests[ip] == RATE_LIMIT {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
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

func JWTAuth(c *gin.Context) {
	if strings.Contains(c.Request.URL.Path, "users") {
		c.Next()
		return
	}
	ibearToken := c.Request.Header.Get("Authorization")
	str_split := strings.Split(ibearToken, " ")
	if len(str_split) == 2 {
		if !jwt.VerifySignature(str_split[1]) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"Message": "Signature is not valid",
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"Message": "Authentication token not provided.",
		})
		return
	}
	c.Next()
}
