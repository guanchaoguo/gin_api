package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/logging"
	"zimuzu_web_api/pkg/util"
	"time"
	"strings"
	"strconv"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		uid := models.GetUidByUserAndPwd(username,password)
		if uid>0 {
			token, err := util.GenerateToken(uid,username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS

				loginTimeTemp:=time.Now().Unix()
				loginTime,_:=strconv.Atoi(strconv.FormatInt(loginTimeTemp,10))
				loginIp:=strings.Split(c.Request.Host,":")[0]
				models.UpdateLastLoginTimeAndIp(uid,loginTime,loginIp)
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
