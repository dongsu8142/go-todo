package models

type Todo struct {
	Id        uint `gorm:"primaryKey"`
	Completed bool
	TodoName  string
}
