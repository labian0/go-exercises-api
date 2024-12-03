package api

import "github.com/gin-gonic/gin"

var EX1_ANSWER int
var API_URL string

func Init(address string, ex1Answer int) {
	API_URL = address
	EX1_ANSWER = ex1Answer
}

func Run() error {
	router := gin.Default()
	endpoints := getDefaultEndpoints()
	for _, ep := range endpoints {
		router.Handle(ep.method, ep.route, ep.endpoint)
	}
	return router.Run(API_URL)
}
