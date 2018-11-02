package api

import (
	"github.com/gin-gonic/gin"
		"zimuzu_web_api/models"
	"net/http"
	)

func ChangePassword(c *gin.Context) {
	username := c.PostForm("username")
	old_password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")

	uid := models.GetUidByUserAndPwd(username,old_password)
	if uid>0 {
		models.UpdatePassword(username,new_password)
		c.JSON(http.StatusOK,gin.H{
			"code":1,
			"msg":"修改成功",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"msg":"修改失败",
		})
	}
}