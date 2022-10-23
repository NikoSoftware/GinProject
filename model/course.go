package model

type Course struct {
	CNo     string  `gorm:"column:c_no"`
	CName   string  `gorm:"column:c_name"`
	TNo     string  `gorm:"column:t_no"`
	Teacher Teacher `gorm:"foreignKey:t_no;references:TNo"`
}

func (Course) TableName() string {
	return "course"
}
