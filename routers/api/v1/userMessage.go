package v1

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
		"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/template"
)

func GetNewUserMsg(c *gin.Context) {
	var dataReturn []models.UserMessage
	template.GetTemplate(c,e.CACHE_PRE_NEWUSERMESSAGE,SetNewUserMsgRedis,dataReturn)
}

func SetNewUserMsgRedis(uid int, redisKey string) {
	newUserMassages:=models.GetNewUserMessage(uid)
	gredis.Set(redisKey,newUserMassages,60*10)
}