package api

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.Default()
	endpoints := getDefaultEndpoints()
	for _, ep := range endpoints {
		router.Handle(ep.method, ep.route, ep.endpoint)
	}
	router.Run("0.0.0.0:1337")
}
