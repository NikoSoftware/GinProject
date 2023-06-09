package model

type Student struct {
	//gorm.Model
	SNo   string `gorm:"column:s_no" json:"sNo"`
	SName string `gorm:"column:s_name" json:"sName"`
	SAge  int    `gorm:"column:s_age" json:"sAge"`
	SSex  string `gorm:"column:s_sex" json:"sSex"`
	Sc    []Sc   `gorm:"foreignKey:s_no;references:SNo"`
}

func (Student) TableName() string {
	return "student"
}
