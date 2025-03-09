package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:eventid", getEvent)
	server.POST("/events", postEvent)
	server.PUT("/events/:eventid", updateEvent)
	server.DELETE("/events/:eventid", deleteEvent)
	server.POST("/signup", signup)
}
