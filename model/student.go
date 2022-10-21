package model

import (
	"time"
)

type Student struct {
	//gorm.Model
	SId   string    `gorm:"column:SId"`
	Sname string    `gorm:"column:Sname"`
	Sage  time.Time `gorm:"column:Sage"`
	Ssex  string    `gorm:"column:Ssex"`
}

func (Student) TableName() string {
	return "Student"
}
