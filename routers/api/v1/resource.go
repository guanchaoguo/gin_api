/**
	资源 相关接口
 */
package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"zimuzu_web_api/models"
	"zimuzu_web_api/pkg/e"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/logging"
)

/**
	获取最新资源接口
 */
func GetHotResourceData(c *gin.Context) {

	// 初始化值
	source := "redis"
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	redisKey := e.CACHE_RESOURCE
	var dataS []models.Resource

	// 先取redis， 没有数据则取mysql存入到redis
	existsTag := gredis.Exists(redisKey)
	if !existsTag {
		dataS = models.GetHotResource(18)
		gredis.Set(redisKey, dataS, e.CACHE_EXPIRE_DAY)
		source = "mysql"
	}

	if len(dataS) <= 0{
		data, err := gredis.Get(redisKey)
		if err != nil {
			code = e.ERROR
			msg = "get redis error,"+err.Error()
			logging.Info(err)
		}

		if len(data) <= 0 {
			code = e.ERROR
			msg = "get redis empty"
		}else {
			errShal := json.Unmarshal(data, &dataS)
			if errShal != nil {
				code = e.ERROR
				msg = "redis data Unmarshal error,"+err.Error()
				logging.Info(err)
			}
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":dataS,
		"source":source,
	})
}

/**
	用户自登录以来 剧集更新的情况
 */
func GetUpgradeInfo(uid int) (re map[string]interface{}) {

	return
}

func ResourceRedisGetData(redisKey string, )  {
	
}


/**
	资源列表 搜索接口
	接口可以接收
		参数有：类别 TypeId、地区 Area、类型 Category、年代PublishYear、电台TvStation、
	    排序：按照最后更新 update_time 按發布日期 create_time 按排名 rank 按評分 rank_value 按點擊率 views
 */
func GetResourceDataByParams(c *gin.Context) {

	// 初始化值
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	resource := make([]models.ResourceReturn, 0)

	// 处理参数
	params, source := _getParams(c)

	page := 1
	if pageStr, isExist := c.GetQuery("Page"); isExist == true {
		pageInt, _ := strconv.Atoi(pageStr)
		page = pageInt
	}

	pageSize := 20
	if pageSizeStr, isExist := c.GetQuery("PageSize"); isExist == true {
		pageSizeInt, _ := strconv.Atoi(pageSizeStr)
		pageSize = pageSizeInt
	}

	// 先取redis 没有取mysql
	from := "redis"
	redisKey := _paramsCreateRedisKey(source, page, pageSize)
	var resourceTmp []models.Resource
	if ! gredis.Exists(redisKey) {
		resourceTmp = models.GetResourceByParams(params, page, pageSize)
		gredis.Set(redisKey, resourceTmp, e.CACHE_EXPIRE_MINUTE)
		from = "mysql"
	}else {

		data, err := gredis.Get(redisKey)
		if err != nil {
			code = e.ERROR
			msg = "get redis error,"+err.Error()
			logging.Info(err)
		}

		if len(data) <= 0 {
			code = e.ERROR
			msg = "get redis empty"
		}else {
			errShal := json.Unmarshal(data, &resourceTmp)
			if errShal != nil {
				code = e.ERROR
				msg = "redis data Unmarshal error,"+err.Error()
				logging.Info(err)
			}
		}
	}

	if len(resourceTmp) > 0 {
		resource = MergeDataToResource(resourceTmp)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":resource,
		"source": from,
	})
}

func _paramsCreateRedisKey(params map[string]string, page int, pageSize int) (redisKey string) {

	keyList := []string{ "Area", "Category", "PublishYear", "TvStation", "TypeId" }
	keyTemp := make([]string,7)
	tempStr := ""
	for v, k := range keyList{
		if _, ok := params[k]; ok {
			tempStr = k+"_"+params[k]
			keyTemp[v] = tempStr
		}else {
			keyTemp[v] = k+"_"
		}
	}
	keyTemp[5] = strconv.Itoa(page)
	keyTemp[6] = strconv.Itoa(pageSize)
	keyTempStr := strings.Join(keyTemp, "_")
	keyTempStr = e.CACHE_RESOURCE_SEARCH + e.Md5Str(keyTempStr)

	return keyTempStr
}

func _getParams(c *gin.Context) ( map[string]interface{}, map[string]string) {

	result := make(map[string]string)
	params := make(map[string]interface{})

	if TypeId, isExist := c.GetQuery("TypeId"); isExist == true {
		result["TypeId"] = TypeId
		TypeIdInt, _ := strconv.Atoi(TypeId)
		params["TypeId"] = TypeIdInt
	}

	if AreaStr, isExist := c.GetQuery("Area"); isExist == true {
		result["Area"] = AreaStr
		params["Area"] = AreaStr
	}

	if CategoryStr, isExist := c.GetQuery("Category"); isExist == true {
		result["Category"] = CategoryStr
		params["Category"] = CategoryStr
	}

	if PublishYear, isExist := c.GetQuery("PublishYear"); isExist == true {
		result["PublishYear"] = PublishYear
		params["PublishYear"] = PublishYear
	}

	if TvStation, isExist := c.GetQuery("TvStation"); isExist == true {
		result["TvStation"] = TvStation
		TvStationInt, _ := strconv.Atoi(TvStation)
		params["TvStation"] = TvStationInt
	}
	return params,result
}

func MergeDataToResource(resource []models.Resource) (returnData []models.ResourceReturn) {

	// 拼凑对应的业务数据
	for _, v := range resource {

		var reReturn models.ResourceReturn
		// Rid,Views,Favorites,Rank int
		reReturn.Rid = v.Rid
		reReturn.Views = v.Views
		reReturn.Favorites = v.Favorites
		reReturn.Rank = v.Rank

		// CnName,EnName,PlayStatus,Poster,PublishYear string
		reReturn.CnName = v.CnName
		reReturn.EnName = v.EnName
		reReturn.PlayStatus = v.PlayStatus
		reReturn.Poster = e.GetImagPath(e.Trim(v.Poster), "m")
		reReturn.PublishYear = strconv.Itoa(v.PublishYear)

		// Area,TvStation,Type,Category,Lang,Remark string
		reReturn.Area = v.Area
		reReturn.TvStation = models.GetTvStation(v.TvStation)
		typeData := models.GetTypeData(v.TypeId)
		reReturn.Type = typeData.Cnname
		reReturn.Category = v.Category
		reReturn.Lang = v.Lang
		reReturn.Remark = e.SubString(v.Remark, 0, 30)

		// PremiereTime,PremiereStatus,UpdateTime string
		scheduleData := GetResourceScheduleData(v.Rid)
		if scheduleData == "" {
			reReturn.PremiereTime = ""
		}else {
			reReturn.PremiereTime = scheduleData + " 首播"
		}
		reReturn.UpdateTime = e.TimeFormatShow(v.UpdateTime, "之前")

		// Score float32
		reReturn.Score = v.Score
		// 汇总数据
		returnData = append(returnData, reReturn)
	}

	return
}

/**
	转换剧集 播放状态
	switch seasonSchedule.Status {
		case 1:
			show = "正在连载第"+strconv.Itoa(seasonSchedule.Season)+"季"
		case 2:
			show = "准备开播第"+strconv.Itoa(seasonSchedule.Season)+"季"
		case 3:
			show = "第"+strconv.Itoa(seasonSchedule.Season)+"季完结"
		case 4:
			show = "全剧完结"
		default:
			show = "未知"
		}
 */
func GetResourceScheduleData(rid int) (returnData string){
	seasonSchedule, err := models.GetResourceLastSchedule(rid)
	fmt.Println(seasonSchedule)
	if err == nil {
		returnData = seasonSchedule.Premiere
	}
	return returnData
}

type DetailReturn struct {
	Base models.ResourceBase
	Score models.ResourceScore
}

/**
	影视详情-基本信息接口
 */
func GetResourceBaseData(c *gin.Context)  {

	// 初始化值
	code := e.SUCCESS
	msg := e.GetMsg(e.SUCCESS)
	var resource DetailReturn

	Rid := 0
	if RidStr, isExist := c.GetQuery("Rid"); isExist == true {
		Rid, _ = strconv.Atoi(RidStr)
	}

	// 基本信息 base
	var _baseDetail models.ResourceBase
	_baseDetail, code = _GetResourceBase(Rid)
	if code == e.SUCCESS {
		resource.Base = _baseDetail
	}

	// 评分信息 score
	var _scoreDetail models.ResourceScore
	_scoreDetail, code = _GetResourceScores(Rid)
	if code == e.SUCCESS {
		resource.Score = _scoreDetail
	}

	// 封面剧照 poster

	// 同类影片 films

	// 相关资讯 news

	// 相关求档 docs

	// 评论信息 comments


	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":msg,
		"data":resource,
	})
}

/**
	基本信息 缓存5分钟
 */
func _GetResourceBase(rid int) (resourceBase models.ResourceBase, errCode int) {

	redisKey := e.CACHE_RESOURCE_DETAIL + strconv.Itoa(rid) + ":base"
	if !gredis.Exists(redisKey) {
		errCode = e.SUCCESS
		resource := models.GetResourceByRid(rid)
		resourceBase.Rid = rid
		resourceBase.Rank = resource.Rank
		resourceBase.CnName = resource.CnName
		resourceBase.EnName = resource.EnName
		resourceBase.AliasName = resource.AliasName
		resourceBase.PlayStatus = resource.PlayStatus
		resourceBase.Poster = e.GetImagPath(resource.Poster, "m")
		resourceBase.PublishYear = strconv.Itoa(resource.PublishYear)
		resourceBase.Area = resource.Area
		resourceBase.TvStation = models.GetTvStation(resource.TvStation)
		typeData := models.GetTypeData(resource.TypeId)
		resourceBase.Type = typeData.Cnname
		resourceBase.Category = resource.Category
		resourceBase.Lang = resource.Lang
		resourceBase.Remark = e.SubString(resource.Remark, 0, 100)
		scheduleData := GetResourceScheduleData(rid)
		if scheduleData != "" {
			resourceBase.PremiereTime = e.StrTimeToWeek(scheduleData)
		}else {
			resourceBase.PremiereTime = "暂无"
		}
		gredis.Set(redisKey, resourceBase, e.CACHE_EXPIRE_MINUTE_5)
	}else {
		data, err := gredis.Get(redisKey)
		if err != nil {
			errCode = e.ERROR_NOT_EXIST_RESOURCE
		}
		if len(data) <= 0 {
			errCode = e.ERROR_NOT_EXIST_RESOURCE
		}else {
			errShal := json.Unmarshal(data, &resourceBase)
			if errShal != nil {
				errCode = e.ERROR_NOT_EXIST_RESOURCE
			}else {
				errCode = e.SUCCESS
			}
		}
	}

	return
}

/**
	评分 缓存5分钟
 */
func _GetResourceScores(rid int) (resourceScore models.ResourceScore, errCode int) {

	redisKey := e.CACHE_RESOURCE_DETAIL + strconv.Itoa(rid) + ":score"
	if !gredis.Exists(redisKey) {
		errCode = e.SUCCESS
		resourceScore = models.GetScoreData(rid)
		resourceScore.FivePre = e.NumberFormat(resourceScore.FiveTotal, resourceScore.Total) * 100
		resourceScore.FourPre = e.NumberFormat(resourceScore.FourTotal, resourceScore.Total) * 100
		resourceScore.ThreePre = e.NumberFormat(resourceScore.ThreeTotal, resourceScore.Total) * 100
		resourceScore.TwoPre = e.NumberFormat(resourceScore.TwoTotal, resourceScore.Total) * 100
		resourceScore.OnePre = e.NumberFormat(resourceScore.OneTotal, resourceScore.Total) * 100
		gredis.Set(redisKey, resourceScore, e.CACHE_EXPIRE_MINUTE_5)
	}else {
		data, err := gredis.Get(redisKey)
		if err != nil {
			errCode = e.ERROR_NOT_EXIST_RESOURCE
		}
		if len(data) <= 0 {
			errCode = e.ERROR_NOT_EXIST_RESOURCE
		}else {
			errShal := json.Unmarshal(data, &resourceScore)
			if errShal != nil {
				errCode = e.ERROR_NOT_EXIST_RESOURCE
			}else {
				errCode = e.SUCCESS
			}
		}
	}

	return
}

