package app

import (
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/pkg/e"
	"net/http"
	"zimuzu_web_api/pkg/logging"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}

func (g *Gin) ResponseError(errCode int,err error,) {

	logging.Warn(err)

	g.C.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  e.GetMsg(errCode),
		"data": nil,
	})

	return
}




