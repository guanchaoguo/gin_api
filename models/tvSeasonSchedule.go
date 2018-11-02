package models

import "errors"

var tssStatus = map[int]string {
	1: "running", // 连载
	2: "start", //开播
	3: "completed", //本季完结
	4: "end", //本剧完结
	5: "unknown", // 未知
}

type TvSeasonSchedule struct {
	Id int
	Rid int
	Season int
	Premiere string
	Status int8
	Sort int16
}

func (TvSeasonSchedule) TableName() string {
	return "tv_season_schedule"
}

func GetResourceLastSchedule(rid int) (tvSeasonSchedule TvSeasonSchedule, err error) {
	err = db.Where(&TvSeasonSchedule{Rid:rid}).Order("Season desc, Sort desc").First(&tvSeasonSchedule).Error
	if err != nil {
		return tvSeasonSchedule, errors.New("查询为空")
	}
	return tvSeasonSchedule, nil
}