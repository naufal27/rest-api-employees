package model

type Employee struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Salary    int64  `json:"salary"`
}
