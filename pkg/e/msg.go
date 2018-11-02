package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_LOFIN:                    "请登录",
	ERROR__INVALID:                 "无效请求",

	ERROR_GET_TAGS_FAIL:           "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:          "统计标签失败",
	ERROR_ADD_TAG_FAIL:            "新增标签失败",
	ERROR_EDIT_TAG_FAIL:           "修改标签失败",
	ERROR_DELETE_TAG_FAIL:         "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:         "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:         "导入标签失败",
	ERROR_ADD_ARTICLE_FAIL:        "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:     "删除文章失败",
	ERROR_EXIST_ARTICLE_FAIL:      "获取资讯失败",
	ERROR_EXIST_COMMENT_FAIL:      "获取评论失败",
	ERROR_EDIT_ARTICLE_FAIL:       "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:      "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:       "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:        "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL: "生成文章海报失败",

	ERROR_REDIS_NOTEXIST:  "redis key not exist",
	ERROR_REDIS_GET:       "get redis error,",
	ERROR_REDIS_EMPTY:     "get redis empty",
	ERROR_REDIS_UNMARSHAL: "redis data Unmarshal error,",

	ERROR_VERIFICATION_MOBILE: "手机验证码不匹配",
	ERROR_VERIFACATIN_EMAIL:   "邮箱验证码不匹配",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
