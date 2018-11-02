package models

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"zimuzu_web_api/pkg/util"
		"errors"
	)

type Article struct {
	*ArticleMain
	*ArticleSection
}

const (
	ArticleTable = "article"
	ERROR_INVALID  = "无效的请求"
)

type ArticleMain struct {
	ID         int    `json:"id"`
	UID        int    `json:"uid"`
	Title      string `json:"title"`
	Intro      string `json:"intro"`
	TypeID     int    `json:"type_id"`
	Poster     string `json:"poster"`
	Views      int    `json:"views"`
	CreateTime int    `json:"create_time"`
	IsComment  int    `json:"is_comment"`

	*RationUser
}

type ArticleSection struct {
	Content    string `json:"content"`
	UpdaterID  int    `json:"updater_id"`
	IsVote     int    `json:"is_vote"`
	Status     int    `json:"status"`
	UpdateTime int    `json:"update_time"`
}

type FindArticle struct {
	Title    string
	Order    string
	Genre    int
	PageNum  int
	PageSize int
	Total    int64
	Fields   []string
}

type RationUser struct {
	Nickname string `json:"nickname"`
	UserId   int    `json:"-"`
	UserPic  string `json:"userpic"`
}

type ResponseArticle struct {
	Total    int64 `json:"total"`
	PageSize int `json:"page_size"`
	Articles []*ArticleMain `json:"articles"`
}

func NewFindArticle() *FindArticle {
	return &FindArticle{}
}

func GetArticle(id int) (*Article, error) {
	DB:= db
	var article Article
	err := DB.First(&article, "id = ? AND status = ? ", id, 1).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func GetArticles(f *FindArticle) ([]*ArticleMain, error) {
	var articles []*ArticleMain
	var count int64
	DB := f.Statement().Debug().Find(&articles).Count(&count)

	if DB.Error != nil && DB.Error != gorm.ErrRecordNotFound {
		return nil, DB.Error
	}

	f.Total = count

	return articles ,nil
}

func UpdateArtViews(c *gin.Context) error {
	DB:= db
	var article Article
	article.ID = com.StrTo(c.Param("id")).MustInt()
	err:= DB.Debug().First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	token := c.GetHeader("Authorization")
	if claims, _ := util.ParseToken(token);claims.Uid != article.UID {
		return errors.New(ERROR_INVALID)
	}

	err = DB.Debug().Model(&article).Update("views",gorm.Expr("views + 1")).Error
	if err != nil  {
		return  err
	}

	return nil
}

func UpdateArticleContent( c *gin.Context) error {
	DB:= db
	var article Article
	article.ID = com.StrTo(c.Param("id")).MustInt()
	err:= DB.Debug().First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	token := c.GetHeader("Authorization")
	if claims, _ := util.ParseToken(token);claims.Uid != article.UID {
		return errors.New(ERROR_INVALID)
	}

	content:= com.StrTo(c.Param("content")).String()
	err = DB.Debug().Model(&article).Update("content",content).Error
	if err != nil  {
		return  err
	}

	return nil
}

func DeleteArticle( c *gin.Context) error {
	DB:= db
	var article Article
	article.ID = com.StrTo(c.Param("id")).MustInt()
	err:= DB.Debug().First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	token := c.GetHeader("Authorization")
	if claims, _ := util.ParseToken(token);claims.Uid != article.UID {
		return errors.New(ERROR_INVALID)
	}

	err = DB.Debug().Delete(&article).Error
	if err != nil  {
		return  err
	}

	return nil
}


func (f *FindArticle) Statement() *gorm.DB {
	DB:= db
	DB = DB.Table(ArticleTable)

	if len(f.Fields) > 0 {
		DB = DB.Select(f.Fields)
	}
	if f.Title != "" {
		DB = DB.Where("title = ?", f.Title)
	}

	if f.Order != "" {
		DB = DB.Order(f.Order, true)
	}

	if f.Genre >= 0 {
		DB = DB.Where("type_id = ?", f.Genre)
	}

	if f.PageNum >= 0 {
		DB = DB.Offset(f.PageNum).Limit(f.PageSize)
	}

	return DB
}

func (f *FindArticle) FormatResponse(articles []*ArticleMain) *ResponseArticle {

	if len(articles) < 1 {
		return nil
	}

	uid := make([]int64, len(articles))
	for i, v := range articles {
		uid[i] = int64(v.UID)
	}

	userInfo := GetUsersByIdIn(uid)
	for i, v := range articles {
		for _, val := range userInfo {
			if v.UID == val.Uid {
				articles[i].Nickname = val.Nickname
				articles[i].UserPic = val.Userpic
			}
		}
	}

	return &ResponseArticle{
		Articles: articles,
		Total:f.Total,
		PageSize:f.PageSize,
	}
}


