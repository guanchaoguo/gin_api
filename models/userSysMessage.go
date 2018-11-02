package models

type UserSysMessage struct {
	Id int `gorm:"private_key;colomn:id"`
	ToUid int `gorm:"colomn:to_uid"`
	Content string `gorm:"colomn:content"`
	ToRead int `gorm:"colomn:to_read"`
	CreateTime int  `gorm:"colomn:create_time"`
}

func (UserSysMessage) TableName() string {
	return "user_sys_message"
}

func GetNewUserSysMessage(uid int) (userSysMessages []UserSysMessage)  {
	db.Where(&UserSysMessage{ToUid:uid,ToRead:0}).Find(&userSysMessages)
	return
}