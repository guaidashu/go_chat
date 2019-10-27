/**
  create by yy on 2019-09-25
*/

package redis

import (
	"github.com/guaidashu/go_helper/redis_tool"
	"go_chat/app/config"
)

var Redis *redis_tool.RedisConn

func InitRedis() (err error) {
	Redis, err = getConnect()
	return
}

func getConnect() (*redis_tool.RedisConn, error) {
	rds := &redis_tool.RedisConn{}
	conf := &redis_tool.RedisConfig{
		Host:     config.Config.Redis.RedisHost,
		Port:     config.Config.Redis.RedisPort,
		Password: config.Config.Redis.RedisPassword,
	}
	err := rds.Init(conf)
	return rds, err
}
