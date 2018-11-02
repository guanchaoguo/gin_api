package v1

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
	"net/http"
	"strconv"
	"zimuzu_web_api/pkg/e"
)

func GetFunctions(c *gin.Context) {
	t,err:=strconv.Atoi(c.Query("type"))
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":e.INVALID_PARAMS,
			"msg":e.GetMsg(e.INVALID_PARAMS),
		})
	} else {
		functions:=models.GetFunctions(t)
		c.JSON(http.StatusOK,gin.H{
			"code":e.SUCCESS,
			"msg":e.GetMsg(e.SUCCESS),
			"data":functions,
		})
	}
}
