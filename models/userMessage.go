package models

type UserMessage struct {
	Id int `gorm:"private_key;colomn:id"`
	FromUid int `gorm:"colomn:from_uid"`
	FromUsername string `gorm:"colomn:from_username"`
	FromDelete int `gorm:"colomn:from_delete"`
	ToUid int `gorm:"colomn:to_uid"`
	ToUsername string `gorm:"colomn:to_username"`
	ToRead int `gorm:"colomn:to_read"`
	ToDelete int `gorm:"colomn:to_delete"`
	Content string `gorm:"colomn:content"`
	Status int `gorm:"colomn:status"`
	CreateTime int `gorm:"colomn:create_time"`
	UpdateTime int `gorm:"colomn:update_time"`
}

func (UserMessage) TableName() string {
	return "user_message"
}

func GetNewUserMessage(uid int) (userMessages []UserMessage) {
	db.Where(&UserMessage{ToUid:uid,ToRead:0}).Find(&userMessages)
	return
}
