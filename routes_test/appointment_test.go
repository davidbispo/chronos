package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"chronos-scheduler.com/api/db"
	"chronos-scheduler.com/api/models"
	"chronos-scheduler.com/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AppointmentTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *AppointmentTestSuite) SetupSuite() {
	db.InitDB()

	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	s.router.POST("/appointments", routes.CreateAppointment)
	s.router.POST("/appointments/attendees", routes.AddAttendeesToAppointment)
	s.router.DELETE("/appointments/attendees", routes.RemoveAttendeesFromAppointment)
}

func (s *AppointmentTestSuite) TestCreateAppointment() {
	appointment := models.Appointment{
		ID:          fmt.Sprintf("appt-%d", time.Now().Unix()),
		Title:       "Test Appointment",
		Description: "Test Description",
		StartTime:   time.Now().Add(1 * time.Hour),
		EndTime:     time.Now().Add(2 * time.Hour),
		Status:      "scheduled",
	}
	body, _ := json.Marshal(appointment)

	req, _ := http.NewRequest("POST", "/appointments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	s.NoError(err)
	s.Equal("created", response["status"])

	var result models.Appointment
	err = db.DB.First(&result, "id = ?", appointment.ID).Error
	s.NoError(err)
	s.Equal(appointment.Title, result.Title)

	db.DB.Delete(&result)
}

func (s *AppointmentTestSuite) TestAddAttendeesToAppointment() {
	appointment := models.Appointment{
		ID:        fmt.Sprintf("appt-%d", time.Now().Unix()),
		Title:     "With Attendee",
		StartTime: time.Now().Add(time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
		Status:    "scheduled",
	}
	attendee := models.Attendee{
		ID:     fmt.Sprintf("att-%d", time.Now().Unix()),
		UserID: 1,
		Name:   "John Doe",
		Email:  "john@example.com",
	}

	db.DB.Create(&appointment)
	db.DB.Create(&attendee)
	defer db.DB.Delete(&appointment)
	defer db.DB.Delete(&attendee)

	links := []models.AppointmentAttendee{
		{AppointmentID: appointment.ID, AttendeeID: attendee.ID},
	}
	body, _ := json.Marshal(links)

	req, _ := http.NewRequest("POST", "/appointments/attendees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)

	var count int
	db.DB.Table("appointments_attendees").
		Where("appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID).
		Count(&count)

	s.Equal(1, count)

	db.DB.Exec("DELETE FROM appointments_attendees WHERE appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID)
}

func (s *AppointmentTestSuite) TestRemoveAttendeesFromAppointment() {
	appointment := models.Appointment{
		ID:        fmt.Sprintf("appt-%d", time.Now().Unix()),
		Title:     "To Remove",
		StartTime: time.Now().Add(time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
		Status:    "scheduled",
	}
	attendee := models.Attendee{
		ID:     fmt.Sprintf("att-%d", time.Now().Unix()),
		UserID: 2,
		Name:   "Jane Doe",
		Email:  "jane@example.com",
	}

	db.DB.Create(&appointment)
	db.DB.Create(&attendee)

	link := models.AppointmentAttendee{
		AppointmentID: appointment.ID,
		AttendeeID:    attendee.ID,
	}
	db.DB.Create(&link)

	defer db.DB.Delete(&appointment)
	defer db.DB.Delete(&attendee)

	body, _ := json.Marshal([]models.AppointmentAttendee{link})
	req, _ := http.NewRequest("DELETE", "/appointments/attendees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)

	var count int
	db.DB.Table("appointments_attendees").
		Where("appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID).
		Count(&count)

	s.Equal(0, count)
}

func TestAppointmentTestSuite(t *testing.T) {
	suite.Run(t, new(AppointmentTestSuite))
}
