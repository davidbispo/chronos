package models

type AppointmentAttendee struct {
	AppointmentID string `json:"appointment_id"`
	AttendeeID    string `json:"attendee_id"`
	Role          string `json:"role,omitempty"`
}
