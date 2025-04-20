package models

type Appointment struct {
	ID          string     `gorm:"primary_key;not null" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartTime   string     `gorm:"not null" json:"start_time"`
	EndTime     string     `gorm:"not null" json:"end_time"`
	Status      string     `gorm:"not null" json:"status"`
	Attendees   []Attendee `gorm:"many2many:appointments_attendees;" json:"attendees"`
}
