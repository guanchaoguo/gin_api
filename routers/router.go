package routers

import (
	"github.com/gin-gonic/gin"

	"zimuzu_web_api/middleware/jwt"
	"zimuzu_web_api/pkg/setting"
	"zimuzu_web_api/routers/api"
	"zimuzu_web_api/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")

	// **********---------- guancg start ---------------- **********
	{
		apiV1.GET("/elastic", v1.GetResources)                 // 默认关键词排行
		apiV1.GET("/elastic/keyword", v1.GetKeywordByResource) // 关键词检索

		apiV1.GET("/articles", v1.GetArticles) //文章列表

		apiV1.GET("/articles/:id", v1.GetArticle)                      //获取文章
		apiV1.PUT("/articles/:id", jwt.JWT(), v1.UpdateArticleContent) //更新文章
		apiV1.DELETE("/articles/:id", jwt.JWT(), v1.DeleteArticle)     //删除文章

		apiV1.GET("/comments/:tid/:id", v1.GetComments)             //评论列表
		apiV1.POST("/comment/:id", jwt.JWT(), v1.StoreComment)      //评论文章
		apiV1.POST("/commentReply/:id", jwt.JWT(), v1.ReplyComment) // 回复评论

		apiV1.POST("/articlesViews/:id", jwt.JWT(), v1.SetArticleViews) //添加文章浏览量

		apiV1.GET("/articlesVote/:id", v1.GetArticlesVotes)            //获取文章投票列表
		apiV1.POST("/articlesVote/:id", jwt.JWT(), v1.SetArticlesVote) //进行投票

	}
	// **********---------- guancg end ---------------- **********

	// ***------------ hukang start--------***
	{
		apiV1.GET("/favTvs", v1.GetFavMovies)
		apiV1.GET("/favAmeTvs", v1.GetFavAmeTvs)
		apiV1.GET("/favJapTvs", v1.GetFavJapTvs)
		// apiV1.GET("/favKorTvs",v1.GetFavKorTvs)
		apiV1.GET("/favOthers", v1.GetFavOthers)
		apiV1.GET("/favArticles", v1.GetFavArticles)
		apiV1.GET("/newUserSysMessages", v1.GetNewUserSysMsg)
		apiV1.GET("/newUserMessages", v1.GetNewUserMsg)
		apiV1.GET("/myInteract", v1.GetMyInteract)
	}
	// ***------------ hukang end--------***

	// ***------------ huzp start--------***
	{
		apiV1.GET("/weixins", v1.GetWeixinShow)
		apiV1.GET("/partners", v1.GetPartnersData)
		apiV1.GET("/resources", v1.GetHotResourceData)
		apiV1.GET("/resourcesParams", v1.GetResourceDataByParams)
		apiV1.GET("/resourceHot", v1.GetHotResourceData)
		apiV1.GET("/resourceBase", v1.GetResourceBaseData)
	}
	// ***------------ huzp end--------***

	apiV1.Use(jwt.JWT())

	return r
}
