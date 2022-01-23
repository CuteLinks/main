package redis

import (
	"cutelinks/config"
	"github.com/go-redis/redis"
	"log"
)

func getRedisConn(dbnum int) *redis.Client{
	return redis.NewClient(&redis.Options{
		Addr: config.GetConfig().REDIS_ADDR,
		Password: "",
		DB: dbnum,
	})
}

func GetUrl(tinyUrl string) string {
	client0:=getRedisConn(0)
	res,err:=client0.Get(tinyUrl).Result()
	if err!=nil{
		log.Println(err.Error())
		return ""
	}
	return res
}
func SaveUrl(tinyUrl string, fullUrl string) string {
	client0:=getRedisConn(0)
	client1:=getRedisConn(1)
	res1,err:=client1.Get(fullUrl).Result()
	if err!=nil{
		log.Println(err.Error())
	}
	if res1!=""{ // если во втором словаре есть value то url не уникально
		return ""
	}
	err1:=client1.Set(fullUrl,"1",0).Err()
	if err1!=nil{
		log.Println(err.Error())
		return ""
	}

	err0:=client0.Set(tinyUrl,fullUrl,0).Err()
	if err0!=nil{
		log.Println(err)
		return ""
	}
	return tinyUrl
}