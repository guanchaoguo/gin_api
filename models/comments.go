package models

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type Comments struct {
   *CommentsMain
   *RationCommentUser
}

const (
	CommentsTable = "comments"
)

type CommentsMain struct {
	Cid          int `json:"cid"`
	UID          int `json:"uid"`
	Good         int `json:"good"`
	Bad          int `json:"bad"`
	Content      string `json:"content"`
	IsMain       int `json:"is_main"`
	UpdateTime   int `json:"update_time"`
}

type CommentsSection struct {
	DocumentID   int `json:"document_id"`
	DocumentType int `json:"document_type"`
	Season       int `json:"season"`
	Episode      int `json:"episode"`
	ReplyID      int `json:"reply_id"`
	ReplyNum     int `json:"reply_num"`
	TopTime      int `json:"top_time"`
	Source       int `json:"source"`
	IP           string `json:"ip"`
	Status       int `json:"status"`
	CreateTime   int `json:"create_time"`
}

type RationCommentUser struct {
	Nickname string `json:"nickname"`
	UserId   int    `json:"-"`
	UserPic  string `json:"userpic"`
}

type FindComments struct {
	Id  int
	Tid int
	PageNum  int
	PageSize int
	Total    int64
}

type CommentsResponse struct {
	Total    int64 `json:"total"`
	PageSize int `json:"page_size"`
	Comments []*Comments `json:"comments"`
}

func NewFindComments()*FindComments{
	return &FindComments{}
}

func GetComments(f *FindComments ) ([]*Comments,error){
	var comments []*Comments
	var count int64
	err := f.CommentsStatement().Debug().Find(&comments).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	f.Total = count

	return comments,err

}

func CreateComment(c *gin.Context) error {

	return nil
}

func CreateReply(c *gin.Context) error {

	return nil
}



func (f *FindComments) CommentsStatement() *gorm.DB {
	DB:= db
	DB = DB.Table(CommentsTable).Where("status = ?", 1)

	if 	 f.Id  > 0 {
		DB = DB.Where("document_id = ?", f.Id)
	}
	if f.Tid  > 0 {
		DB = DB.Where("document_type = ?", f.Tid)
	}

	if f.PageNum >= 0 {
		DB = DB.Offset(f.PageNum).Limit(f.PageSize)
	}

	DB = DB.Order("top_time desc").Order("update_time desc")

	return DB
}


func   (f *FindComments) CommentsFormat(comments []*Comments) *CommentsResponse{
	if len(comments) < 1 {
		return nil
	}

	uid := make([]int64, len(comments))
	for i, v := range comments {
		uid[i] = int64(v.UID)
	}

	userInfo := GetUsersByIdIn(uid)
	for i, v := range comments {
		for _, val := range userInfo {
			if v.UID == val.Uid {
				comments[i].Nickname = val.Nickname
				comments[i].UserPic = val.Userpic
			}
		}
	}

	 return &CommentsResponse{
		Total:    f.Total,
		PageSize: f.PageSize,
		Comments: comments,
	}
}