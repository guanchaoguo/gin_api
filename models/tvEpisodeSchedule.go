package models

type TvEpisodeSchedule struct {
	Id int
	Rid int
	Season int8
	Episode int8
	PlayTime string
	Status int8
	Sort int16
}

func (TvEpisodeSchedule) TableName() string {
	return "tv_episode_schedule"
}

func GetResourceLastEpisode(rid int) (tvEpisodeSchedule []TvEpisodeSchedule) {

	db.Where(&TvEpisodeSchedule{Rid:rid}).Order("Season desc, Episode desc, Sort desc").First(&tvEpisodeSchedule)
	return
}
