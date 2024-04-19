package api

import "github.com/gin-gonic/gin"

func GetRouters() *gin.Engine {
	router := gin.Default()

	router.POST("/message", postMessage)
	router.POST("/voice", postVoice)

	return router
}
