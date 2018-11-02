package v1

import (
	"net/http"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/app"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/util"
	"zimuzu_web_api/pkg/setting"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	article, err := models.GetArticle(id)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EXIST_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	f := models.NewFindArticle()

	f.Order = ""
	if arg := c.PostForm("sort"); arg != "" {
		f.Order = com.StrTo(arg).String()
		valid.Required(f.Order, "sort").Message("排序条件不能为空")
	}

	f.Genre = -1
	if arg := c.PostForm("genre"); arg != "" {
		f.Genre = com.StrTo(arg).MustInt()
		valid.Min(f.Genre, 0, "genre").Message("类型必须大于0")
	}

	f.Title = ""
	if arg := c.PostForm("title"); arg != "" {
		f.Title = com.StrTo(arg).String()
		valid.Required(f.Title, "title").Message("查询条件不能为空")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	f.Fields = []string{"id","poster","intro","title","uid","views","create_time",}
	f.PageNum = util.GetPage(c)
	f.PageSize = setting.PageSize

	articles, err := models.GetArticles(f)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, f.FormatResponse(articles))
}


func UpdateArticleContent(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	content := com.StrTo(c.Param("content")).String()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.Required(content, "content").Message("修改内容不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.UpdateArticleContent(c); err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteArticle(c *gin.Context){
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if err := models.DeleteArticle(c); err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}


func SetArticleViews(c *gin.Context)  {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}


	if err := models.UpdateArtViews(c); err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}



