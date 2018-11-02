package models

import (
	"fmt"
)

type Resource struct {
	Rid int `gorm:"column:rid"`
	CnName string `gorm:"column:cnname"`
	EnName string `gorm:"column:enname"`
	AliasName string `gorm:"column:aliasname"`
	Remark string `gorm:"column:remark"`
	TypeId int `gorm:"column:type_id"`
	Category string `gorm:"column:category"`
	Area string `gorm:"column:area"`
	Lang string `gorm:"column:lang"`
	Format string `gorm:"column:format"`
	PublishYear int `gorm:"column:publish_year"`
	TvStation int `gorm:"column:tvstation"`
	PlayStatus string `gorm:"column:play_status"`
	Zimuzu string `gorm:"column:zimuzu"`
	Rank int `gorm:"column:rank"`
	RankValue float32 `gorm:"column:rank_value"`
	Poster string `gorm:"column:poster"`
	PremiereTime int `gorm:"column:premiere_time"`
	PlayTime int `gorm:"column:play_time"`
	Views int `gorm:"column:views"`
	Favorites int `gorm:"column:favorites"`
	Score float32 `gorm:"column:score"`
	Operator int `gorm:"column:operator"`
	UpdateTime int64 `gorm:"column:update_time"`
	CreateTime int64 `gorm:"column:create_time"`
	Status int `gorm:"column:status"`
}

func (Resource) TableName() string {
	return "resource"
}

func GetResourceAreaFromId(id int) string {
	var resource Resource
	db.Where(&Resource{Rid:id}).First(&resource)
	return resource.Area
}

func GetHotResource(hot int) (resource []Resource){
	db.Where(&Resource{Status:1}).Order("rank asc", true).Limit(hot).Find(&resource)
	return
}

func GetResourceByRid(Rid int) (resource Resource){
	db.Where(&Resource{Rid:Rid}).First(&resource)
	return
}

type ResourceReturn struct {
	Rid,Views,Favorites,Rank int
	CnName,EnName,PlayStatus,Poster,PublishYear string
	Area,TvStation,Type,Category,Lang,Remark string
	PremiereTime,PremiereStatus,UpdateTime string
	Score float32
}

func GetResourceByParams(params interface{}, page int, pageSize int) (resource []Resource) {

	// 数据库
	DB := db

	// 解析参数
	temp := params.(map[string]interface{})
	for kk, value := range temp{
		if kk == "Category" {
			switch tempType := value.(type) {
			case string:
				DB = DB.Where("Category like ?", "%"+tempType+"%")
			default:
				fmt.Println("unknown Category: ", value)
			}
		}else if kk == "TypeId"{
			switch tempType := value.(type) {
			case string:
				DB = DB.Where("type_id in (?)", []string{"", "0", tempType})
			case int:
				DB = DB.Where("type_id in (?)", []int{0, tempType})
			default:
				fmt.Println("unknown TypeId: ", value)
			}
		}else if kk == "Area"{
			switch tempType := value.(type) {
			case string:
				DB = DB.Where("area in (?)", []string{"", tempType})
			default:
				fmt.Println("unknown Area: ", value)
			}
		}else if kk == "PublishYear"{
			switch tempType := value.(type) {
			case int:
				DB = DB.Where("publish_year in (?)", []int{0, tempType})
			case int64:
				DB = DB.Where("publish_year in (?)", []int64{0, tempType})
			case string:
				DB = DB.Where("publish_year in (?)", []string{"0", tempType})
			default:
				fmt.Println("unknown PublishYear: ", value)
			}
		}else if kk == "TvStation"{
			switch tempType := value.(type) {
			case int:
				DB = DB.Where("tvstation in (?)", []int{0, tempType})
			default:
				fmt.Println("unknown TvStation: ", value)
			}
		}
	}

	// 取 resource 数据库 数据
	DB.Limit(pageSize).Offset((page - 1) * pageSize).Find(&resource)
	return
}