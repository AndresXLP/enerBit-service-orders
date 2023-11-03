package redis

import (
	"fmt"

	"enerBit-service-orders/config"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
)

type redisOptions struct {
	host string
	port int
}

func (r *redisOptions) getAddress() string {
	return fmt.Sprintf("%s:%d", r.host, r.port)
}

func NewConnection() redis.Conn {
	address := redisOptions{
		host: config.Environments().Redis.Host,
		port: config.Environments().Redis.Port,
	}

	connection, err := redis.Dial("tcp", address.getAddress())
	if err != nil {
		panic(err)
	}

	connection.Do("XGROUP", "CREATE", "mystream", "mygroup", "$", "MKSTREAM")

	log.Info("Redis Stream Connection Successfully")
	return connection
}
