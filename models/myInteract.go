package models

type MyInteract struct {
	Id int `gorm:"private_key;colomn:id"`
	Uid int `gorm:"colomn:uid"`
	NewComment string `gorm:"colomn:new_comment"`
	NewReply string `gorm:"colomn:new_reply"`
	NewGood string `gorm:"colomn:new_good"`
	NewBad string `gorm:"colomn:new_bad"`
	LastOpenTime string `gorm:"colomn:last_open_time"`
}

func (MyInteract) TableName() string {
	return "my_interact"
}

func GetMyInteract(uid int) (myInteract MyInteract) {
	db.Where(&MyInteract{Uid:uid}).First(&myInteract)
	return
}
