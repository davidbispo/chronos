package models

type AppointmentAttendee struct {
	AppointmentID string `json:"appointment_id" gorm:"primary_key;not null"`
	AttendeeID    string `json:"attendee_id" gorm:"primary_key;not null"`
	Role          string `json:"role,omitempty"`
}

func (AppointmentAttendee) TableName() string {
	return "appointments_attendees"
}
