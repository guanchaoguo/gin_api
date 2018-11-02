package models

type Partner struct {
	Id int `gorm:"column:id"`
	Name string `gorm:"column:name"`
	Pic string `gorm:"column:pic"`
	Desc string `gorm:"column:desc"`
	Url string `gorm:"column:url"`
	Sort int `gorm:"column:sort"`
	Status int `gorm:"column:status"`
	Operator int `gorm:"column:operator"`
	CreateTime int `gorm:"column:create_time"`
}

func (Partner) TableName() string {
	return "partner"
}

func GetPartners(params interface{}) (partners []Partner) {
	db.Where(params).Order("sort asc", true).Limit(12).Find(&partners)
	return
}

func GetPartnersTest() (partners []Partner) {
	db.Where("Status = ?", 1).Where("Url like ?", "%www%").Where("Sort > ?", 0).Order("sort asc", true).Limit(4).Find(&partners)
	return
}