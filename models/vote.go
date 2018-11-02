package models

import "github.com/jinzhu/gorm"

type Vote struct {
	ID        int `json:"id"`
	Title     string `json:"title"`
	MultCheck     string `json:"mult_check"`
	CheckNum     string `json:"check_num"`
	OptionId    int `json:"-"`
	VoteId    int `json:"-"`
    Option       string `json:"-"`
	Num int `json:"-"`
	Percent  int `json:"-"`

	VoteOption []*VoteOption `json:"vote_option"`
}
type VoteOption struct {
	ID      int `json:"id"`
	Option  string `json:"option"`
	Num     int `json:"num"`
	Percent int `json:"percent"`
}

func GetArticlesVote( id int)([]Vote,error){
	DB:= db
	var votes []Vote
	err :=DB.Debug().Table("vote").
		Select("vote.id, vote.title ,vote.mult_check,vote.check_num ,o.id as option_id,o.vote_id,o.option,o.num ,o.percent").
		Where(" aid = ?", id).
		Joins("inner join vote_option o ON o.vote_id = vote.id").Find(&votes).Error

	if err != nil  {
		return nil ,err
	}

	tmp:=  make(map[int]int)
	var voteList  []Vote
	var i int
	for _,v:= range votes {
		option := voteOption(v)
		if key, ok := tmp[v.VoteId]; ok {
			voteList[key].VoteOption = append(voteList[key].VoteOption ,option)
		}else{
			v.VoteOption = append(v.VoteOption ,option)
			voteList = append(voteList ,v)
			tmp[v.VoteId] = i
			i ++
		}
	}

	return voteList ,nil
}

func SetArticlesVote(id int)(bool,error) {
	DB:= db
	var vote VoteOption
	vote.ID = id
	err:= DB.Debug().First(&vote).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false ,err
	}

	err = DB.Debug().Model(&vote).Update("num",gorm.Expr("num + 1")).Error
	if err != nil  {
		return false ,err
	}

	return true,nil
}

func voteOption(v Vote)* VoteOption {
	return &VoteOption{
		ID:v.VoteId,
		Option:v.Option,
		Num : v.Num,
		Percent : v.Percent,
	}
}