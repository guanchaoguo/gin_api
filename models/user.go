package models

import (
	"fmt"
)
type User struct {
	Uid           int    `gorm:"private_key;colomn:uid"`
	Sex           int    `gorm:"colomn:sex"`
	Email         string `gorm:"colomn:email"`
	Password      string `gorm:"colomn:password"`
	Salt          string `gorm:"colomn:salt"`
	Nickname      string `gorm:"colomn:nickname"`
	PhoneArea     int    `gorm:"colomn:phone_area"`
	PhoneNumber   string `gorm:"colomn:phone_number"`
	Signature     string `gorm:"colomn:signature"`
	Userpic       string `gorm:"colomn:userpic"`
	MainGroupId   int    `gorm:"colomn:main_group_id"`
	GroupExpire   int    `gorm:"colomn:group_expire"`
	Level         int    `gorm:"colomn:level"`
	Point         int    `gorm:"colomn:point"`
	LastLoginTime int    `gorm:"colomn:last_login_time"`
	LastLoginIp   string `gorm:"colomn:last_login_ip"`
	LoginCount    int    `gorm:"colomn:login_count"`
	UpdateTime    int    `gorm:"colomn:update_time"`
	CreateTime    int    `gorm:"colomn:create_time"`
	CreateIp      string `gorm:"colomn:create_ip"`
	Status        int    `gorm:"colomn:status"`
	Client        int    `gorm:"colomn:client"`
	IosDid        string `gorm:"colomn:ios_did"`
	JpushId       string `gorm:"colomn:jpush_id"`
	AndriodDid    string `gorm:"colomn:andriod_did"`
}

func (User) TableName() string {
	return "user"
}

func GetUserLastLoginTime(uid int) int {
	var user User
	db.Where(&User{Uid:uid}).First(&user)
	return user.LastLoginTime
}


func GetUsersByIdIn(uid []int64) []*User {
	var users []*User
	DB:= db
	DB.Debug().Table("user").Select("uid ,userpic,nickname").Where("uid in (?)",uid).Find(&users).Limit(15)
	return users
}

func GetUidByUserAndPwd(nickname,password string) int {
	var user User
	db.Where(User{Nickname:nickname,Password:password}).First(&user)
	return user.Uid
}

func UpdateLastLoginTimeAndIp(uid, lastLoginTime int,lastLoginIp string) {
	var user User
	db.Where(User{Uid:uid}).First(&user)
	db.Model(&user).Updates(User{
		LastLoginTime:lastLoginTime,
		LastLoginIp:lastLoginIp,
	})
}

func AddUser(nickname,password string)  {
	user:=User{Nickname:nickname,Password:password}
	fmt.Println(db.Create(&user).Error)
}

func UpdatePassword(nickname string, password string) {
	var user User
	db.Where(User{Nickname:nickname}).First(&user)
	db.Model(&user).Updates(User{
		Password:password,
	})
}

