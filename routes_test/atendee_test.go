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

type AttendeeTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *AttendeeTestSuite) SetupSuite() {
	db.InitDB()

	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	s.router.POST("/attendees", routes.CreateAttendee)
}

func (s *AttendeeTestSuite) TestCreateAttendee() {
	attendee := models.Attendee{
		ID:       fmt.Sprintf("att-%d", time.Now().Unix()),
		UserID:   100,
		Name:     "Alice Tester",
		Email:    "alice@example.com",
		Metadata: `{"source":"test"}`,
	}

	body, _ := json.Marshal(attendee)
	req, _ := http.NewRequest("POST", "/attendees", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	s.NoError(err)
	s.Equal("created", response["status"])

	var result models.Attendee
	err = db.DB.First(&result, "id = ?", attendee.ID).Error
	s.NoError(err)
	s.Equal(attendee.Email, result.Email)

	db.DB.Delete(&result)
}

func TestAttendeeTestSuite(t *testing.T) {
	suite.Run(t, new(AttendeeTestSuite))
}
