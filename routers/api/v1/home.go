package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/logging"
)

func GetYesterdayHotDownload(c *gin.Context) {
	// 初始化值
	source := "redis"
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	redisKey := e.CACHE_HOME_RESOURCE_HOT
	var dataS []models.Resource

	// 先取redis， 没有数据则取mysql存入到redis
	existsTag := gredis.Exists(redisKey)
	if !existsTag {
		dataS = models.GetHotResource(18)
		gredis.Set(redisKey, dataS, e.CACHE_EXPIRE_DAY)
		source = "mysql"
	}

	if len(dataS) <= 0{
		data, err := gredis.Get(redisKey)
		if err != nil {
			code = e.ERROR
			msg = "get redis error,"+err.Error()
			logging.Info(err)
		}

		if len(data) <= 0 {
			code = e.ERROR
			msg = "get redis empty"
		}else {
			errShal := json.Unmarshal(data, &dataS)
			if errShal != nil {
				code = e.ERROR
				msg = "redis data Unmarshal error,"+err.Error()
				logging.Info(err)
			}
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":dataS,
		"source":source,
	})
	
}
