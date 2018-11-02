package v1

import (
"github.com/gin-gonic/gin"
	"net/http"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/logging"
	)


//获取关键词前十的影视
func GetResources (c *gin.Context) {
	data, err := gredis.ZRangeByScore("myzset")
	if err != nil {
		logging.Info(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": data,
	})
}

//获取关键词检索
func GetKeywordByResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  e.GetMsg(200),
		"data": "",
	})

}