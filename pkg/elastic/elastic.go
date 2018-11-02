package elastic
//
//import (
//	"zimuzu_web_api/pkg/setting"
//	"log"
//	"time"
//	"os"
//	"github.com/olivere/elastic"
//)
//
//var ElasticClient *elastic.Client
//
//func init() {
//	var err error
//
//
//	sec, err := setting.Cfg.GetSection("elastic")
//	if err != nil {
//		log.Fatal(2, "Fail to get section 'database': %v", err)
//	}
//
//	URL := sec.Key("Password").String()
//
//
//	ElasticClient, err := elastic.NewClient(
//		elastic.SetURL(URL),
//		elastic.SetSniff(false),
//		elastic.SetHealthcheckInterval(10*time.Second),
//		elastic.SetRetrier(NewCustomRetrier()),
//		elastic.SetGzip(true),
//		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
//		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
//
//}
