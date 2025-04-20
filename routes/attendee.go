package routes

import (
	"net/http"

	"chronos-scheduler.com/api/db"
	"chronos-scheduler.com/api/models"
	"github.com/gin-gonic/gin"
)

func CreateAttendee(c *gin.Context) {
	var a models.Attendee
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&a).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attendee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "created"})
}
