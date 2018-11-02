/**
	友情链接 相关接口
 */
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

func GetPartnersData(c *gin.Context) {

	// 初始化值
	source := "redis"
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	redisKey := e.CACHE_PARTNERS
	var dataS []models.Partner

	// 先取redis， 没有数据则取mysql存入到redis
	existsTag := gredis.Exists(redisKey)
	if !existsTag {
		params := make(map[string]interface{})
		params["Status"] = 1
		// 测试多条件查询
		//if c.Query("ID") != "" {
		//	params["Id"] = c.Query("ID")
		//}
		//if c.Query("Operator") != "" {
		//	params["Operator"] = c.Query("Operator")
		//}
		dataS = models.GetPartners(params)
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