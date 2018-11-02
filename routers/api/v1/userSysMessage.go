package v1

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
		"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/template"
)

func GetNewUserSysMsg(c *gin.Context) {
	var dataReturn []models.UserMessage
	template.GetTemplate(c,e.CACHE_PRE_NEWUSERSYSMESSAGE,setNewUserSysMsgRedis,dataReturn)
}

func setNewUserSysMsgRedis(uid int, redisKey string) {
	newUserSysMassages:=models.GetNewUserSysMessage(uid)
	gredis.Set(redisKey,newUserSysMassages,60*10)
}