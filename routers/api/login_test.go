package api

import (
	"testing"
	"strings"
	"time"
	"strconv"
)

func TestHostSplit(t *testing.T) {
	host1:="192.168.3.1:80"
	host2:="192.168.3.2"
	ip1:=strings.Split(host1,":")[0]
	ip2:=strings.Split(host2,":")[0]
	t.Log(ip1,ip2)

	ti:=time.Now()
	t.Log(ti)
	t.Log(ti.Unix())
	t.Log(ti.String())
	t.Log(ti.UnixNano())

	var v int64=9999999999
	s:=strconv.FormatInt(v,10)
	i,_:=strconv.Atoi(s)
	t.Log(i)
}
