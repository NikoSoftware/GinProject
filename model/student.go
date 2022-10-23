package model

type Student struct {
	//gorm.Model
	SNo   string `gorm:"column:s_no"`
	SName string `gorm:"column:s_name"`
	SAge  int    `gorm:"column:s_age"`
	SSex  string `gorm:"column:s_sex"`
	Sc    []Sc   `gorm:"foreignKey:s_no;references:SNo"`
}

func (Student) TableName() string {
	return "student"
}
