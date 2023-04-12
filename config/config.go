package config

import (
	"rest-api-employee/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:L3zOXUDH92OIUry6vJDg@tcp(containers-us-west-206.railway.app:6111)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connection DB")
	}
	migration()
}
func migration() {
	DB.AutoMigrate(&model.Employee{})
}
