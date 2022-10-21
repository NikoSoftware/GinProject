package model

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	SId   string
	Sname string
	Sage  time.Time
	Ssex  string
}

func (s Student) TableName() string {
	return "student"
}
