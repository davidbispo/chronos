package main

import (
	"net/http"

	"chronos-scheduler.com/api/config"
	"chronos-scheduler.com/api/db"
	"chronos-scheduler.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	r.POST("/attendees", func(c *gin.Context) {
		routes.CreateAttendee(c)
	})

	r.POST("/appointments", func(c *gin.Context) {
		routes.CreateAppointment(c)
	})

	r.POST("/appointments/attendees", func(c *gin.Context) {
		routes.AddAttendeesToAppointment(c)
	})

	r.DELETE("/appointments/attendees", func(c *gin.Context) {
		routes.RemoveAttendeesFromAppointment(c)
	})

	r.Run() // defaults to :8080
}
