package models

type DailyDownloads struct {
	Id int `gorm:"private_key;colomn:id"`
	ItemId int `gorm:"colomn:item_id"`
	Counts int `gorm:"colomn:counts"`
	Time int `gorm:"colomn:time"`
}

func (DailyDownloads) TableName() string {
	return "daily_downloads"
}

func GetDailyHotDownloads(time int, num int) (dailyDownloads []DailyDownloads)  {
	db.Debug().Where(&DailyDownloads{Time:time}).Order("counts desc").Limit(num).Find(&dailyDownloads)
	return
}


