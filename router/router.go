package router

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(c *gin.Context) {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8081",
	})
	proxy.ServeHTTP(c.Writer, c.Request)
}

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/calculator.sum", ReverseProxy)
	router.POST("/calculator.sub", ReverseProxy)
	router.POST("/calculator.mul", ReverseProxy)
	router.POST("/calculator.div", ReverseProxy)

	return router
}
