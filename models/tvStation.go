package models

type TvStation struct {
	Id int `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (TvStation) TableName() string {
	return "tvstation"
}

func GetTvStation(id int) (name string) {

	var tvstation TvStation
	db.Where(&TvStation{Id:id}).First(&tvstation)
	if len(tvstation.Name) > 0 {
		name = tvstation.Name
	}
	return
}

