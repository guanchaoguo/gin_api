package v1

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/template"
)

func GetMyInteract(c *gin.Context) {
	var dataReturn models.MyInteract
	template.GetTemplate(c,e.CACHE_PRE_NEWMYINTERACT,setMyInteractRedis,dataReturn)
}

func setMyInteractRedis(uid int, redisKey string)  {
	myInteract:=models.GetMyInteract(uid)
	gredis.Set(redisKey,myInteract,60*10)
}