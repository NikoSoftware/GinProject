package model

type Sc struct {
	SNo    string `gorm:"column:s_no"`
	CNo    string `gorm:"column:c_no"`
	Score  string `gorm:"column:score"`
	Course Course `gorm:"foreignKey:c_no;references:CNo"`
}

func (Sc) TableName() string {
	return "sc"
}
