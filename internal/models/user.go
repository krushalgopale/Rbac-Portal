package models

type User struct {
	ID       uint   `gorm:"primaryKey;autoincrement"`
	Name     string `                                json:"name"`
	Email    string `gorm:"unique"                   json:"email"`
	Password string `                                json:"password"`
	Role     string `                                json:"role"`
}
