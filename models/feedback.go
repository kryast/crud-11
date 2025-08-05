package models

type Feedback struct {
	ID           uint   `gorm:"type:primary_key" json:"id"`
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
	Message      string `json:"message"`
	Rating       uint   `json:"rating"`
}
