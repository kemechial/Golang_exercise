package routes

import (
	"example.com/testserver/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var ShutdownChan = make(chan struct{})


func getEvent(context *gin.Context) {

	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.Save()
	context.JSON(http.StatusCreated, event)

}

func shutdownServer(context *gin.Context) {	
	context.JSON(http.StatusOK, gin.H{"message": "Server is shutting down..."})
	go func() {
		time.Sleep(100 * time.Millisecond) // Let response flush
		ShutdownChan <- struct{}{}
	}()
}

func updateEvent(context *gin.Context) {	
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}	

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Could not fetch the event"})
		return
	}

	var updateEvent models.Event
	if err := context.ShouldBindJSON(&updateEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.ID = eventID
	updateEvent.Update()
	context.JSON(http.StatusOK, updateEvent)
}

func deleteEvent(context *gin.Context){
	id := context.Param("id")
	eventID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
