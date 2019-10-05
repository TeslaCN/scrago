package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

const (
	host = "127.0.0.1"
	port = 6379
)

func TestRedis(t *testing.T) {
	conn, e := redis.Dial("tcp", ":6379")
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
