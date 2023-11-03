package repo

import (
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
)

type Repository interface {
	SendStreamLog(message interface{})
}

type repository struct {
	streamRedis redis.Conn
}

func NewRedisRepository(streamRedis redis.Conn) Repository {
	return &repository{streamRedis}
}

func (r repository) SendStreamLog(message interface{}) {
	if _, err := r.streamRedis.Do("XADD", "mystream", "*", "message", message); err != nil {
		log.Error(err)
	}
}
