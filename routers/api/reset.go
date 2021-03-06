package api

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/pkg/vcode"
	"zimuzu_web_api/models"
	"net/http"
	"zimuzu_web_api/pkg/e"
)

func Reset(c *gin.Context) {
	mobile:=c.PostForm("mobile")
	mobileVcode:=c.PostForm("vcode")
	nickname:=c.PostForm("nickname")
	password:=c.PostForm("password")
	errCode:=vcode.VeriMobileVcode(mobile,mobileVcode)
	if errCode==0 {
		models.UpdatePassword(nickname,password)
		c.JSON(http.StatusOK,gin.H{
			"code":errCode,
			"msg":"修改成功",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code":errCode,
			"msg":e.GetMsg(errCode),
		})
	}

}