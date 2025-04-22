package routes

import (
	"net/http"

	"chronos-scheduler.com/api/db"
	"chronos-scheduler.com/api/models"
	"github.com/gin-gonic/gin"
)

func CreateAppointment(c *gin.Context) {
	var a models.Appointment
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&a).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

func AddAttendeesToAppointment(c *gin.Context) {
	var links []models.AppointmentAttendee
	if err := c.BindJSON(&links); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(links) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No attendees to add"})
		return
	}

	for _, link := range links {
		if err := db.DB.Create(&link).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":          "Failed to add attendee",
				"attendee_id":    link.AttendeeID,
				"appointment_id": link.AppointmentID,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "attendees added"})
}

func RemoveAttendeesFromAppointment(c *gin.Context) {
	var links []models.AppointmentAttendee
	if err := c.BindJSON(&links); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, link := range links {
		if err := db.DB.Delete(&link).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove attendee", "id": link.AttendeeID})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "attendees removed"})
}
