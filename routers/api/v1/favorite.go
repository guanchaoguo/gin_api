package v1

import (
	"github.com/gin-gonic/gin"
		"zimuzu_web_api/models"
	"net/http"
	"zimuzu_web_api/pkg/e"
	. "github.com/ahmetb/go-linq"
	"strconv"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/logging"
	"encoding/json"
)

func setFavoriteRedis(uid int) {
	favorites:=models.GetFavorites(uid)

	var articleFavs,tvFavs,movieFavs,openclassFavs []models.Favorite
	From(favorites).Where(func(f interface{}) bool {
		return f.(models.Favorite).ChannelId==1
	}).ToSlice(&articleFavs)
	From(favorites).Where(func(f interface{}) bool {
		return f.(models.Favorite).ChannelId==2 && f.(models.Favorite).TypeId==20
	}).ToSlice(&tvFavs)
	From(favorites).Where(func(f interface{}) bool {
		return f.(models.Favorite).ChannelId==2 && f.(models.Favorite).TypeId==21
	}).ToSlice(&movieFavs)
	From(favorites).Where(func(f interface{}) bool {
		return f.(models.Favorite).ChannelId==2 && f.(models.Favorite).TypeId==22
	}).ToSlice(&openclassFavs)

	var tvAmeFavs,tvJapFavs,tvKorFavs,otherFavs []models.Favorite
	for _, tvFav := range tvFavs {
		area:=models.GetFavoriteTvArea(tvFav)
		switch area {
		case "美国":
			tvAmeFavs=append(tvAmeFavs, tvFav)
			break
		case "日本":
			tvJapFavs=append(tvJapFavs,tvFav)
			break
		case "韩国":
			tvKorFavs=append(tvKorFavs,tvFav)
			break
		default:
			otherFavs=append(otherFavs,tvFav)
			break
		}
	}
	otherFavs=append(otherFavs,openclassFavs...)

	uidStr:=strconv.Itoa(uid)
	gredis.Set(e.CACHE_PRE_FAVORITE_MOVIE+uidStr,movieFavs,60*10)
	gredis.Set(e.CACHE_PRE_FAVORITE_TV_AMERICA+uidStr,tvAmeFavs,60*10)
	gredis.Set(e.CACHE_PRE_FAVORITE_TV_JAPAN+uidStr,tvJapFavs,60*10)
	gredis.Set(e.CACHE_PRE_FAVORITE_TV_KOREA+uidStr,tvKorFavs,60*10)
	gredis.Set(e.CACHE_PRE_FAVORITE_OTHER+uidStr,otherFavs,60*10)
	gredis.Set(e.CACHE_PRE_FAVORITE_ARTICLE+uidStr,articleFavs,60*10)
}

func shareGet(c *gin.Context,redisPreKey string)  {
	c.Set("uid",1)
	if uid, ok := c.Get("uid"); ok {

		code := e.SUCCESS
		msg := e.GetMsg(e.SUCCESS)
		var dataS []models.Favorite

		redisKey:=redisPreKey+strconv.Itoa(uid.(int))

		existTag:=gredis.Exists(redisKey)
		if !existTag {
			setFavoriteRedis(uid.(int))
		}
		data,err:=gredis.Get(redisKey)
		if err != nil {
			code = e.ERROR
			msg = "get redis error,"+err.Error()
			logging.Info(err)
		}
		if string(data)=="null" {
			code=e.ERROR
			msg="get redis empty"
		}else {
			errShal:=json.Unmarshal(data,&dataS)
			if errShal!=nil {
				code=e.ERROR
				msg = "redis data Unmarshal error,"+errShal.Error()
				logging.Info(errShal)
			}
		}

		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"msg":msg,
			"data":dataS,
		})
	}
}

func GetFavMovies(c *gin.Context) {
	shareGet(c,e.CACHE_PRE_FAVORITE_MOVIE)
}

func GetFavAmeTvs(c *gin.Context) {
	shareGet(c,e.CACHE_PRE_FAVORITE_TV_AMERICA)
}

func GetFavJapTvs(c *gin.Context) {
	shareGet(c,e.CACHE_PRE_FAVORITE_TV_JAPAN)
}

func GetFavKorTvs(c *gin.Context) {
	shareGet(c,e.CACHE_PRE_FAVORITE_TV_KOREA)
}

func GetFavOthers(c *gin.Context) {
	shareGet(c,e.CACHE_PRE_FAVORITE_OTHER)
}

func GetFavArticles(c *gin.Context) {
	shareGet(c,e.CACHE_PRE_FAVORITE_ARTICLE)
}

