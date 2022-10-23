package model

type Teacher struct {
	TNo   string `gorm:"column:t_no;primaryKey"`
	TName string `gorm:"column:t_name"`
}

func (Teacher) TableName() string {
	return "teacher"
}
