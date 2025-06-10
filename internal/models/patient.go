package models

import (
	"time"
)

type Patient struct {
	ID        uint        `gorm:"primaryKey;autoincrement"`
	Name      string      `                                json:"name"`
	Email     string      `gorm:"unique"                   json:"email"`
	Phone     int         `                                json:"phone"`
	Age       int         `                                json:"age"`
	Gender    string      `                                json:"gender"`
	Disease   string      `                                json:"disease"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
