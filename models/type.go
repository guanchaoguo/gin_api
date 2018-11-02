package models

type Type struct {
	ID       int
	ChannelId string
	Cnname string
	Enname string
	PageTitle string
	PageKeywords string
	PageDescription string
	Status int
	UpdateTime int
	CreateTime int
}

func (Type) TableName() string {
	return "type"
}

func GetTypeData(typeId int) (types Type) {
	db.Where(Type{Status: 1,ID:typeId}).First(&types)
	return
}
