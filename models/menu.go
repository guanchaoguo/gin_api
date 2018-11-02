package models
//菜单
type Menu struct {
	Id uint `gorm:"primary_key;column:Id"`
	Name string `gorm:"column:name"`
}

func (Menu) TableName() string {
	return "menus"
}

func GetMenus() (menus []Menu) {
	db.Find(&menus)
	return
}

//Logo
type Logo struct {
	Id  uint `gorm:"primary_key;column:Id"`
	Img []byte `gorm:"column:img"`
}

func (Logo) TableName() string {
	return "logo"
}

func GetLogo() (logo Logo) {
	db.First(&logo)
	return
}

//APP下载链接
type AppInfo struct {
	Id          uint `gorm:"primary_key;column:Id"`
	Img         []byte `gorm:"column:img"`
	Android_url string `gorm:"column:android_url"`
	Iphone_url  string `gorm:"column:iphone_url"`
	H5_url      string `gorm:"column:h5_url"`
}

func (AppInfo) TableName() string {
	return "app_info"
}

func GetAppInfo() (appInfo AppInfo) {
	db.First(&appInfo)
	return
}