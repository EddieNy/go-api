package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:eventid", getEvent)
	server.POST("/events", postEvent)
	server.PUT("/event/:eventid", updateEvent)
}
