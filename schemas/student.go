package schemas

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string
	Email   string
	Age     int
	Course  string
	Address string
	Phone   string
}
