package models

type AppointmentAttendee struct {
	AppointmentID string `json:"appointment_id" gorm:"primary_key;not null;index:idx_appointment_attendee"`
	AttendeeID    string `json:"attendee_id" gorm:"primary_key;not null;index:idx_appointment_attendee"`
	Role          string `json:"role,omitempty"`
	RSVPStatus    string `json:"rsvp_status" gorm:"primary_key;not null;index"`
}

func (AppointmentAttendee) TableName() string {
	return "appointments_attendees"
}
