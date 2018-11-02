package v1

import (
	"encoding/json"
	"strconv"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/logging"
)

func GetType(typeId int) (typeData models.Type) {

	redisKey := e.CACHE_TYPE+":"+strconv.Itoa(typeId)
	existsTag := gredis.Exists(redisKey)
	if !existsTag {
		typeData = models.GetTypeData(typeId)
		gredis.Set(redisKey, typeData, e.CACHE_EXPIRE_MONTH)
	}else {
		getData, err := gredis.Get(redisKey)
		if err != nil {
			logging.Info(err)
		}
		errShal := json.Unmarshal(getData, &typeData)
		if errShal != nil {
			logging.Info(errShal.Error())
		}
	}
	return typeData
}
