package models

type Attendee struct {
	ID       int64  `gorm:"primary_key;not null;autoIncrement" json:"id"`
	UserID   int64  `gorm:"type:bigint;not null" json:"user_id"`
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null" json:"email"`
	Metadata string `json:"metadata"`
}
