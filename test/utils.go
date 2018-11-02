package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
)

var (
	router    *gin.Engine
	myHeaders = []CustomizedHeader{}
	logging   *log.Logger
)

type CustomizedHeader struct {
	Key     string
	Value   string
	IsValid bool
}

func SetRouter(r *gin.Engine) {
	router = r
}

func SetLog(l *log.Logger) {
	logging = l
}

func printfLog(format string, v ...interface{}) {
	if logging == nil {
		return
	}

	logging.Printf(format, v...)
}

func changeToFieldName(name string) string {
	result := ""
	i := 0
	j := 0
	r := []rune(name)
	for m, v := range r {
		if v >= 'A' && v < 'a' {
			if (m != 0 && r[m-1] >= 'a') || ((m != 0 && r[m-1] >= 'A' && r[m-1] < 'a') && (m != len(r)-1 && r[m+1] >= 'a')) {
				i = j
				j = m
				result += name[i:j] + "_"
			}
		}
	}

	result += name[j:]
	return strings.ToLower(result)
}

func getQueryStr(params interface{}) (result string, err error) {
	if params == nil {
		return
	}
	value := reflect.ValueOf(params)

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			result += "&" + changeToFieldName(value.Type().Field(i).Name) + "=" + fmt.Sprintf("%v", value.Field(i).Interface())
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			result += "&" + fmt.Sprintf("%v", key.Interface()) + "=" + fmt.Sprintf("%v", value.MapIndex(key).Interface())
		}
	default:
		err = ErrMustBeStructOrMap
		return
	}

	if result != "" {
		result = result[1:]
	}
	return
}

func runHandler(req *http.Request) (bodyByte []byte, err error) {

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	bodyByte, err = ioutil.ReadAll(result.Body)
	return
}

func Requested(method string, uri string, way string, param interface{}) (bodyByte []byte, err error) {
	if router == nil {
		err = ErrRouterNotSet
		return
	}

	var req *http.Request
	switch way {
	case Json:
		jsonBytes := []byte{}
		if param != nil {
			jsonBytes, err = json.Marshal(param)
			if err != nil {
				return
			}
		}
		req = httptest.NewRequest(method, uri, bytes.NewReader(jsonBytes))
		req.Header.Set("Content-Type", "application/json")

		printfLog("TestOrdinaryHandler\tRequest:\t%v:%v,\trequestBody:%v\n", method, uri, string(jsonBytes))
	case Form:
		queryStr := ""
		if param != nil {
			queryStr, err = getQueryStr(param)
			if err != nil {
				return
			}
		}

		if queryStr != "" {
			queryStr = "?" + queryStr
		}
		req = httptest.NewRequest(method, uri+queryStr, nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		printfLog("TestOrdinaryHandler\tRequest:\t%v:%v%v\n", method, uri, queryStr)
	}

	if myHeaders != nil {
		for _, data := range myHeaders {
			if data.IsValid {
				req.Header.Set(data.Key, data.Value)
			}
		}
	}

	bodyByte, err = runHandler(req)

	printfLog("TestOrdinaryHandler\tResponse:\t%v:%v\tResponse:%v\n", method, uri, string(bodyByte))
	return
}

func UnMarshalResp(method string, uri string, way string, param interface{}, resp interface{}) error {
	bodyByte, err := Requested(method, uri, way, param)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyByte, resp)
}
