package routes

import (
	"example.com/rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:eventid", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", postEvent)
	authenticated.PUT("/events/:eventid", updateEvent)
	authenticated.DELETE("/events/:eventid", deleteEvent)
	authenticated.POST("/events/:eventid/register", registerForEvent)
	authenticated.DELETE("/events/:eventid/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
