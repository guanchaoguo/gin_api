package gredis

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"

	"zimuzu_web_api/pkg/setting"
	"log"
	"time"
	"zimuzu_web_api/models"
	"strings"
	"strconv"
)

var RedisConn *redis.Pool

func init() {
	var err error


	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	Password := sec.Key("Password").String()
	Host := sec.Key("Host").String()
	MaxIdle ,err:= sec.Key("MaxIdle").Int()
	MaxActive,err := sec.Key("MaxActive").Int()
	IdleTimeout,err := sec.Key("IdleTimeout").Int64()

	RedisConn = &redis.Pool{
		MaxIdle:    MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout:  time.Duration(IdleTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", Host)
			if err != nil {
				return nil, err
			}
			if Password != "" {
				if _, err := c.Do("AUTH", Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}

func Set(key string, data interface{}, time int) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	reply, err := redis.Bool(conn.Do("SET", key, value))
	conn.Do("EXPIRE", key, time)

	return reply, err
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func ZScore(key ,item string )(int, error)  {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ZSCORE", key,item))
	if err != nil {

		return 0, err
	}
	return reply, nil
}

func Zincrby(key ,item string)(int, error)  {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Int(conn.Do("ZINCRBY", 1,item))
	if err != nil {

		return 0, err
	}
	return reply, nil
	
}

func ZRangeByScore(key string)([]models.Search,error)  {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Values(conn.Do("ZREVRANGEBYSCORE", key, "+inf","-9" ))
	if err != nil {
		return nil, err
	}

	var list []models.Search
	for _, v := range reply {
		str := strings.Split(string(v.([]byte)), ":")
		id,err := strconv.Atoi(str[0])
		if err != nil {
			return nil, err
		}

		tid,err := strconv.Atoi(str[2])
		if err != nil {
			return nil, err
		}

		list = append(list,models.Search{Id:id,Name: str[2],Type:tid})
	}
	return list, nil
}