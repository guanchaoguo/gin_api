package models

import "testing"

func TestUpdateLastLoginTimeAndIp(t *testing.T) {
	UpdateLastLoginTimeAndIp(1,1540277700,"192.168.3.123")
}

func TestGetUidByUserAndPwd(t *testing.T) {
	uid:=GetUidByUserAndPwd("小萝莉","81a42f57d108a7b684ed12a31cd529e3")
	t.Log(uid)
}

func TestAddUser(t *testing.T) {
	AddUser("test2","testpwd2")
}