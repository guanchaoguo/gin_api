package v1

import (
	"net/http"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/app"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/setting"
	"zimuzu_web_api/pkg/util"
)

func GetComments(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	f:= models.NewFindComments()

	if arg := c.Param("id"); arg != ""  {
		f.Id = com.StrTo(arg).MustInt()
		valid.Required(f.Id, "title").Message("ID必须大于0")
	}

	if arg := c.Param("tid"); arg != ""  {
		f.Tid = com.StrTo(arg).MustInt()
		valid.Required(f.Tid, "title").Message("TID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	f.PageSize = setting.PageSize
	f.PageNum = util.GetPage(c)
	comments, err := models.GetComments(f)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EXIST_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, f.CommentsFormat(comments))
}

func StoreComment(c *gin.Context){

	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	content := com.StrTo(c.Param("content")).String()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(content, "content").Message("评论内容不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.CreateComment(c); err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func ReplyComment(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	content := com.StrTo(c.Param("content")).String()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(content, "content").Message("回复内容不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.CreateReply(c);err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}




