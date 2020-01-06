package redis

import (
	"fmt"
	"github.com/TeslaCN/scrago/core/setting"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

const (
	host = "rpi4"
	port = 6379
)

func TestRedis(t *testing.T) {
	conn, e := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if e != nil {
		t.Fatal(e)
	}
	reply, _ := conn.Do("set", "redigo", "hello, redigo")
	log.Printf("Reply: %s", reply)
	reply, _ = conn.Do("get", "redigo")
	log.Printf("Reply: %s", reply)
	reply, _ = conn.Do("del", "redigo")
	log.Printf("Reply: %d", reply)
	reply, _ = conn.Do("get", "redigo")
	log.Printf("Reply: %s", reply)
}

func TestInitBitmap(t *testing.T) {
	conn, e := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port), redis.DialPassword(""))
	if e != nil {
		t.Fatal(e)
	}
	_, _ = conn.Do("del", "scrago:dedup")
	re, _ := conn.Do("setbit", "scrago:dedup", setting.GetBloomFilterSize()-1, 0)
	log.Println(re)
	reply, _ := conn.Do("strlen", "scrago:dedup")
	log.Println(reply)
}
