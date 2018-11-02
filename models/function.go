package models

type Function struct {
	Id int `gorm:"primary_key;column:id"`
	Cnname string `gorm:"column:cnname"`
	Enname string `gorm:"column:enname"`
	Type int `gorm:"column:type"`
}

func (Function) TableName() string {
	return "function"
}

func GetFunctions(t int) (functions []Function) {
	db.Where(&Function{Type:t}).Find(&functions)
	return
}
