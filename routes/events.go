package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "could not fetch events, please try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventid"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func postEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not create event"})
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventid"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "could not fetch the event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not update data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event was successfully updated!"})
}
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventid"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch the event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event successfully deleted!"})
}
