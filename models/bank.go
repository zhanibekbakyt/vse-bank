package models

type Bank struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
