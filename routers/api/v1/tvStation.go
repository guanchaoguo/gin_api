package v1

import (
	"strconv"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
)

func GetTvStationById(id int) (name string) {

	redisKey := e.CACHE_TV_STATION+":"+strconv.Itoa(id)
	existsTag := gredis.Exists(redisKey)
	if !existsTag {
		name = models.GetTvStation(id)
		if len(name) > 0 {
			gredis.Set(redisKey, name, e.CACHE_EXPIRE_MONTH)
		}
	}else {
		getData, err := gredis.Get(redisKey)
		if err == nil {
			name = string(getData[:])
		}
	}

	return name
}
