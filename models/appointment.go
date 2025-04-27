package models

import "time"

type Appointment struct {
	ID          int64      `gorm:"primary_key;not null;autoIncrement" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartTime   time.Time  `gorm:"not null" json:"start_time"`
	EndTime     time.Time  `gorm:"not null" json:"end_time"`
	Status      string     `gorm:"not null" json:"status"`
	Attendees   []Attendee `gorm:"many2many:appointments_attendees;" json:"attendees"`
}
