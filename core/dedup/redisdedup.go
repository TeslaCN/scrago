package dedup

import (
	"fmt"
	"github.com/TeslaCN/scrago/core/setting"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/url"
	"reflect"
	"time"
)

func init() {
	setting.AddDeduplicationType("RedisDeduplicate", reflect.TypeOf(&RedisDeduplicate{}))
}

type RedisDeduplicate struct {
	pool *redis.Pool
	key  string
}

func NewRedisDeduplicate(host string, port int, password string, key string) Deduplicate {
	rd := &RedisDeduplicate{
		pool: &redis.Pool{
			Dial: func() (conn redis.Conn, err error) {
				return redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port), redis.DialPassword(password))
			},
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
		},
		key: key,
	}
	return rd
}

func (rd *RedisDeduplicate) De(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := HashCode(s)
	position := Reserve(hashCode, setting.GetDeduplicationOffset())

	conn := rd.pool.Get()
	defer conn.Close()
	exists, e := conn.Do("setbit", rd.key, position, 1)
	if e != nil {
		log.Fatalln(e)
	}
	value, ok := exists.(int64)
	if !ok {
		log.Printf("Redis return unexpected type [%t] value [%v]", exists, exists)
		return 0
	}
	return int(value)
}

func (rd *RedisDeduplicate) Exist(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := HashCode(s)
	position := Reserve(hashCode, setting.GetDeduplicationOffset())

	conn := rd.pool.Get()
	defer conn.Close()
	exists, e := conn.Do("getbit", rd.key, position)
	if e != nil {
		log.Println(e)
	}
	value, ok := exists.(int64)
	if !ok {
		log.Printf("Redis return unexpected type [%t] value [%v]", exists, exists)
		return 0
	}
	return int(value)
}

func (rd *RedisDeduplicate) Remove(u url.URL) int {
	s := fmt.Sprintf("%s%s", u.Host, u.RequestURI())
	hashCode := HashCode(s)
	position := Reserve(hashCode, setting.GetDeduplicationOffset())

	conn := rd.pool.Get()
	defer conn.Close()
	reply, e := conn.Do("setbit", rd.key, position, 0)
	if e != nil {
		log.Println(e)
	}
	value, ok := reply.(int64)
	if !ok {
		log.Printf("Redis return unexpected type [%t] value [%v]", reply, reply)
		return 0
	}
	return int(value)
}
