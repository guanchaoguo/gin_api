package vcode

import (
	"math/rand"
	"time"
	"fmt"
	ypclnt	"github.com/yunpian/yunpian-go-sdk/sdk"
	"zimuzu_web_api/pkg/setting"
	"zimuzu_web_api/pkg/gredis"
	"zimuzu_web_api/pkg/e"
	"encoding/json"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sendgrid/sendgrid-go"
)

var clnt=ypclnt.New(setting.MobileKey)

func SendMobileVcode(mobile string) (err error) {
	rnd:=rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode:=fmt.Sprintf("%06v",rnd.Int31n(1000000))

	params:=ypclnt.NewParam(2)
	params[ypclnt.MOBILE]=mobile
	params[ypclnt.TEXT]=fmt.Sprintf("【人人影视】您的验证码是%s",vcode)

	r:=clnt.Sms().SingleSend(params)
	if r.IsSucc() {
		gredis.Set(e.CACHE_PRE_VERIFICATION_MOBILE+mobile,vcode,60)
	}else {
		r.Error(err)
	}
	return
}

func VeriMobileVcode(mobile string, vcodeClient string) (errCode int) {
	redisKey:=e.CACHE_PRE_VERIFICATION_MOBILE+mobile
	if !gredis.Exists(redisKey) {
		errCode=e.ERROR_REDIS_NOTEXIST
		return
	}
	data,erro:=gredis.Get(redisKey)
	if erro != nil {
		errCode=e.ERROR_REDIS_GET
	}
	if string(data) == "null" {
		errCode=e.ERROR_REDIS_EMPTY
	}else {
		var vcode string
		errShal:=json.Unmarshal(data,&vcode)
		if errShal!=nil {
			errCode=e.ERROR_REDIS_UNMARSHAL
		}else {
			if vcodeClient!=vcode {
				errCode=e.ERROR_VERIFICATION_MOBILE
			}
		}
	}
	return
}

var mailCli=sendgrid.NewSendClient(setting.EmailKey)

func SendEmailVcode(emailAddr string) (err error) {
	rnd:=rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode:=fmt.Sprintf("%06v",rnd.Int31n(1000000))

	from:=mail.NewEmail("noreply","noreply@mails.zimuzu.tv")
	to:=mail.NewEmail("h",emailAddr)
	subject:="邮箱验证码"
	context:="您的邮箱验证码为"+vcode
	msg:=mail.NewSingleEmail(from,subject,to,context,context)
	_,err=mailCli.Send(msg)
	if err == nil {
		gredis.Set(e.CACHE_PRE_VERIFICATION_EMAIL+emailAddr,vcode,60)
	}
	return
}

func VeriEmailVcode(emailAddr string, vcodeClient string) (errCode int) {
	redisKey:=e.CACHE_PRE_VERIFICATION_EMAIL+emailAddr
	if !gredis.Exists(redisKey) {
		errCode=e.ERROR_REDIS_NOTEXIST
		return
	}
	data,erro:=gredis.Get(redisKey)
	if erro != nil {
		errCode=e.ERROR_REDIS_GET
	}
	if string(data) == "null" {
		errCode=e.ERROR_REDIS_EMPTY
	}else {
		var vcode string
		errShal:=json.Unmarshal(data,&vcode)
		if errShal!=nil {
			errCode=e.ERROR_REDIS_UNMARSHAL
		}else {
			if vcodeClient!=vcode {
				errCode=e.ERROR_VERIFICATION_MOBILE
			}
		}
	}
	return
}