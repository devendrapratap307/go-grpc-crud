package models

type User struct {
	ID    int64 `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}
