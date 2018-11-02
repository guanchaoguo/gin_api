package search

import (
	"gopkg.in/olivere/elastic.v6"
	"zimuzu_web_api/pkg/logging"
	"zimuzu_web_api/pkg/setting"
)

var Client *elastic.Client

func NewElasticClient() (*elastic.Client, error) {
	if Client != nil {
		return Client, nil
	}
	var err error

	sec, err := setting.Cfg.GetSection("elastic")
	if err != nil {
		logging.Warn(err)
		return nil, err
	}

	host := sec.Key("URL").String()
	Client, err = elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		logging.Warn(err)
		return nil, err
	}
	return Client, nil
}
