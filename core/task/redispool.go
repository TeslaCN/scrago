package task

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/url"
	"strings"
	"time"
)

type RedisPool struct {
	pool *redis.Pool
	key  string
}

// Redis任务池
func NewRedisPool(host string, port int, password string, key string) Pool {
	return &RedisPool{
		pool: &redis.Pool{
			Dial: func() (conn redis.Conn, err error) {
				return redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port), redis.DialPassword(password))
			},
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
		},
		key: key,
	}
}

// Fetch 拉取任务
func (rp *RedisPool) Fetch() (*Task, error) {
	conn := rp.pool.Get()
	defer conn.Close()
	reply, err := conn.Do("brpop", rp.key, 60)
	if err != nil {
		log.Println(err)
	}
	s, ok := reply.([]interface{})
	if !ok {
		log.Println("Fetched %v", reply)
	}
	split := strings.Split(string(s[1].([]byte)), " ")
	parse, err := url.Parse(split[1])
	if err != nil {
		return nil, err
	}
	return &Task{
		Method: split[0],
		Url:    *parse,
	}, nil
}

// Offer 任务入列
func (rp *RedisPool) Offer(task Task) {
	conn := rp.pool.Get()
	defer conn.Close()
	target := fmt.Sprintf("%s %s", task.Method, task.Url.String())
	_, err := conn.Do("lpush", rp.key, target)
	if err != nil {
		log.Println(err)
	}
}
