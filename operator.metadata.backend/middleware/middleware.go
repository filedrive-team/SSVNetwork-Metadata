package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ssv_operator_metadata/controllers"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,NetType")
		if http.MethodOptions == c.Request.Method {
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{})
		} else {
			c.Next()
		}
	}
}

func NetType() gin.HandlerFunc {
	return func(c *gin.Context) {
		netType := c.GetHeader("NetType")
		switch netType {
		case "mainnet", "testnet":
		default:
			c.Abort()
			controllers.JSONError(c, "lose header:"+"NetType")
			return
		}
	}
}

// limit
func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
