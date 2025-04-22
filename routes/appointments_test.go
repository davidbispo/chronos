package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"chronos-scheduler.com/api/db"
	"chronos-scheduler.com/api/models"
	"chronos-scheduler.com/api/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	db.InitDB()

	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.POST("/appointments", routes.CreateAppointment)
	router.POST("/appointments/attendees", routes.AddAttendeesToAppointment)
	router.DELETE("/appointments/attendees", routes.RemoveAttendeesFromAppointment) // ✅ added

	os.Exit(m.Run())
}

func TestCreateAppointment(t *testing.T) {
	appointment := models.Appointment{
		ID:          fmt.Sprintf("appt-%d", time.Now().UnixNano()),
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
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", rec.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	if response["status"] != "created" {
		t.Fatalf("Expected status 'created', got %v", response["status"])
	}

	var result models.Appointment
	err := db.DB.First(&result, "id = ?", appointment.ID).Error
	if err != nil {
		t.Fatalf("Appointment not found in DB: %v", err)
	}

	defer db.DB.Delete(&result)

	if result.Title != appointment.Title {
		t.Fatalf("Expected title '%s', got '%s'", appointment.Title, result.Title)
	}
}

func TestAddAttendeesToAppointment(t *testing.T) {
	appointment := models.Appointment{
		ID:        fmt.Sprintf("appt-%d", time.Now().UnixNano()),
		Title:     "Test",
		StartTime: time.Now().Add(time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
		Status:    "scheduled",
	}
	attendee := models.Attendee{
		ID:     fmt.Sprintf("att-%d", time.Now().UnixNano()),
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
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", rec.Code)
	}

	var count int
	db.DB.Table("appointments_attendees").
		Where("appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID).
		Count(&count)

	if count != 1 {
		t.Fatalf("Expected attendee to be added to appointment, found %d", count)
	}

	db.DB.Exec("DELETE FROM appointments_attendees WHERE appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID)
}

func TestRemoveAttendeesFromAppointment(t *testing.T) {
	appointment := models.Appointment{
		ID:        fmt.Sprintf("appt-%d", time.Now().UnixNano()),
		Title:     "To Remove",
		StartTime: time.Now().Add(time.Hour),
		EndTime:   time.Now().Add(2 * time.Hour),
		Status:    "scheduled",
	}
	attendee := models.Attendee{
		ID:     fmt.Sprintf("att-%d", time.Now().UnixNano()),
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
	defer db.DB.Exec("DELETE FROM appointments_attendees WHERE appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID)

	body, _ := json.Marshal([]models.AppointmentAttendee{link})

	req, _ := http.NewRequest("DELETE", "/appointments/attendees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", rec.Code)
	}

	var count int
	db.DB.Table("appointments_attendees").
		Where("appointment_id = ? AND attendee_id = ?", appointment.ID, attendee.ID).
		Count(&count)

	if count != 0 {
		t.Fatalf("Expected attendee to be removed from appointment, found %d", count)
	}
}
