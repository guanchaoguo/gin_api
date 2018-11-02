package models

/**
	剧情基本信息返回
 */
type ResourceBase struct {
	Rid,Rank int
	CnName,EnName,AliasName,PlayStatus,Poster,PublishYear string
	Area,TvStation,Type,Category,Lang,Remark string
	PremiereTime string
}
/**
	剧情详情
 */
type ResourceDetail struct {
	Id int
	Rid int
	Content string
	PageTitle string
	PageKeywords string
	PageDescription string
	Permission string
	Imdb string
	CloseResource int
	CloseTime int64
	CreateTime int64
}

func GetResourceDetail(rid int) (resourceDetail ResourceDetail) {
	db.Table("resource_detail").Where(ResourceDetail{Rid:rid}).First(&resourceDetail)
	return
}

/**
	主演 导演 编剧
 */
type ResourceCharacter struct {
	Id int
	Rid int
	profession int
	Cnname string
	Enname string
	PageDescription string
	Sort int
	UpdateTime int64
	CreateTime int64
}

func GetResourceCharacter(rid int) (resourceCharacters []ResourceCharacter) {
	db.Table("resource_character").Where(ResourceCharacter{Rid:rid}).Find(&resourceCharacters)
	return
}

/**
	评分表
 */
type Evaluate struct {
	Id int
	Rid int
	Uid int
	Score int
	CmtId int
	CreateTime int64
}

/**
	评分返回值 结构体
 */
type ResourceScore struct {
	Rid int
	Average float64
	Total int
	FiveTotal int
	FivePre float64
	FourTotal int
	FourPre float64
	ThreeTotal int
	ThreePre float64
	TwoTotal int
	TwoPre float64
	OneTotal int
	OnePre float64

}

func GetScoreData(rid int) (resourceScore ResourceScore) {

	resourceScore.Rid = rid
	// Total sum Average
	sql := " COUNT(1) as Total, FORMAT(AVG(score), 1) as Average "
	rows, err := db.Table("evaluate").Where(Evaluate{Rid:rid}).Select(sql).Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var Total int
			var Average float64
			rows.Scan(&Total, &Average)
			resourceScore.Total = Total
			resourceScore.Average = Average
		}
	}

	resourceScore.FiveTotal = _GetEveryData(9, 10)
	resourceScore.FourTotal = _GetEveryData(7, 8)
	resourceScore.ThreeTotal = _GetEveryData(5, 6)
	resourceScore.TwoTotal = _GetEveryData(3, 4)
	resourceScore.OneTotal = _GetEveryData(0, 2)

	return
}

func _GetEveryData(min int, max int) (result int) {

	rows, err := db.Table("evaluate").Where(" score BETWEEN ? AND ? ", min, max).Select(" COUNT(1) as Total ").Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var Total int
			rows.Scan(&Total)
			result = Total
		}
	}
	return
}