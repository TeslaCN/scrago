package dedup

import (
	"log"
	"net/url"
	"testing"
)

func TestRedisDeduplication(t *testing.T) {
	d := NewRedisDeduplicate("rpi4", 6379, "", "scrago:test")
	u, _ := url.Parse("http://rpi3:8080")
	log.Println(d.Exist(*u))
	log.Println(d.De(*u))
	log.Println(d.Exist(*u))
	log.Println(d.Remove(*u))
	log.Println(d.Exist(*u))
}
