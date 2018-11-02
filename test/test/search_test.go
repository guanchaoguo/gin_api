package test

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testing"
	"zimuzu_web_api/routers/api/v1"
	utils "zimuzu_web_api/test"
)

var (
	search = keywords{Keyword: "功夫"}
)

type keywords struct {
	Keyword string `json:"keyword"`
}

type OrdinaryResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func init() {

	router := gin.Default()

	router.GET("/api/v1/elastic", v1.GetResources)
	router.GET("/api/v1/elastic/keyword", v1.GetKeywordByResource)

	utils.SetRouter(router)

	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	utils.SetLog(newLog)
}

func TestGetResources(t *testing.T) {

	resp := OrdinaryResponse{}
	err := utils.UnMarshalResp("GET", "/api/v1/elastic", "form", nil, &resp)
	if err != nil {
		t.Errorf("TestGetResourcesHandler: %v\n", err)
		return
	}
	if resp.Code != 200 {
		t.Errorf("TestGetResourcesHandler: response is not expected\n")
		return
	}
}

func TestGetKeywordByResource(t *testing.T) {

	resp := OrdinaryResponse{}
	err := utils.UnMarshalResp("GET", "/api/v1/elastic/keyword", "form", search, &resp)
	if err != nil {
		t.Errorf("TestGetResourcesHandler: %v\n", err)
		return
	}
	if resp.Code != 200 {
		t.Errorf("TestGetResourcesHandler: response is not expected\n")
		return
	}
}
