package models

type Weixin struct {
	Id int `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
	GroupName string `gorm:"column:group_name"`
	Pic string `gorm:"column:pic"`
	Sort int `gorm:"column:sort"`
	Status int `gorm:"column:status"`
	CreateTime int `gorm:"column:create_time"`
	UpdateTime int `gorm:"column:update_time"`
}

func (Weixin) TableName() string {
	return "weixin_show"
}

func GetShowData(t int) (weixinShows []Weixin) {
	db.Where(&Weixin{Status:t}).Order("sort asc", true).Limit(4).Find(&weixinShows)
	return
}