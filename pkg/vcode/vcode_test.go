package vcode

import "testing"

func TestSentMobileVcode(t *testing.T) {
	err:=SendMobileVcode("****")
	if err!=nil {
		t.Error(err)
	}else {
		t.Log("发送成功")
	}
}

func TestVeriMobileVcode(t *testing.T) {
	err:=VeriMobileVcode("****","387213")
	if err!=nil {
		t.Error(err)
	}else {
		t.Log("验证成功")
	}
}

