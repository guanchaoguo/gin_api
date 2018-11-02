package models

type Favorite struct {
	Id int `gorm:"private_key;colomn:id"`
	ChannelId int `gorm:"colomn:channel_id"`
	TypeId int `gorm:"colomn:type_id"`
	ItemId int `gorm:"colomn:item_id"`
	Uid int `gorm:"colomn:uid"`
	CreateTime int `gorm:"colomn:create_time"`
}

func (Favorite) TableName() string {
	return "favorite"
}

func GetFavorites(uid int) (favorites []Favorite) {
	db.Where(&Favorite{Uid:uid}).Find(&favorites)
	return
}

func GetFavoriteTvArea(favorite Favorite) string {
	return GetResourceAreaFromId(favorite.ItemId)
}
