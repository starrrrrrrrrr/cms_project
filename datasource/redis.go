package datasource

import (
	"cms_project/config"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

//返回redis实例
func NewRedis() *redis.Database {
	var database *redis.Database
	//获取配置
	conf := config.InitConfig()
	if conf != nil {
		rd := conf.Redis
		iris.New().Logger().Info(rd)
		//添加配置
		database = redis.New(redis.Config{
			Network:   rd.NetWork,
			Addr:      rd.Addr + ":" + rd.Port,
			Database:  "",
			MaxActive: 10,
			Timeout:   redis.DefaultRedisTimeout,
			Prefix:    rd.Prefix,
		})
	} else {
		iris.New().Logger().Info("NewRedis error")
	}
	return database
}
