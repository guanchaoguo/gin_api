package template

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/pkg/e"
	"strconv"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/logging"
	"encoding/json"
	"net/http"
)

func GetTemplate(c *gin.Context, redisPreKey string, fSetRedis func(uid int, redisKey string),dataReturn interface{})  {
	c.Set("uid",1)
	if uid, ok := c.Get("uid"); ok {
		code:=e.SUCCESS
		msg:=e.GetMsg(e.SUCCESS)

		redisKey:=redisPreKey+strconv.Itoa(uid.(int))
		if !gredis.Exists(redisKey) {
			fSetRedis(uid.(int),redisKey)
		}

		data,err:=gredis.Get(redisKey)
		if err != nil {
			code = e.ERROR
			msg = "get redis error,"+err.Error()
			logging.Info(err)
		}
		if string(data) == "null" {
			code=e.ERROR
			msg="get redis empty"
		}else {
			errShal:=json.Unmarshal(data,&dataReturn)
			if errShal!=nil {
				code=e.ERROR
				msg = "redis data Unmarshal error,"+errShal.Error()
				logging.Info(errShal)
			}
		}

		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":msg,
			"data":dataReturn,
		})
	}
}
