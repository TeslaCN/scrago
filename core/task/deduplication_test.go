package task

import (
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestUrlParse(t *testing.T) {
	u, _ := url.Parse("https://root:123456@www.baidu.com:8000/login?name=xiaoming&name=xiaoqing&age=24&age1=23#fffffff")
	fmt.Println(u)
	u.RequestURI()
	u.EscapedPath()
	u.Query()
	u.Port()
	u.Hostname()
	//values, _ := url.ParseQuery(u.RawQuery)
	//u2, _ := url.Parse("mysql:https://root:123456@www.baidu.com:0000/login?name=xiaoming&name=xiaoqing&age=24&age1=23#fffffff")
	//fmt.Println(u2)
	fmt.Println()
}

func TestDedup(t *testing.T) {
	deduplicate := NewDeduplicate()
	u0, _ := url.Parse("https://ss9874.com/un/001.htm")
	u1, _ := url.Parse("https://ss9874.com/un/002.htm")
	t.Log(deduplicate.De(*u0))
	t.Log(deduplicate.De(*u1))
	t.Log(deduplicate.De(*u1))
}

var m = make(map[string]string)

func TestMap(t *testing.T) {
	getMap()["hello"] = "world"
	log.Println(m)
}

func getMap() map[string]string {
	return m
}
