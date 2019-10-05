package pipeline

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	path := "d:/cache/x/hnd00674jp-7.jpg"
	fmt.Println(path)
	// file, e := os.Open(path)
	// if e != nil {
	// 	t.Fatal(e)
	// }
	// info, e := os.Stat(path)
	info, e := os.Stat("/")
	if e != nil {
		t.Fatal(e)
	}
	t.Log(info.Name())
	t.Log(info.IsDir())
	t.Log(info.Size())
	infos, e := ioutil.ReadDir("/../d/cache")
	if e != nil {
		t.Fatal(e)
	}
	for _, v := range infos {
		t.Log(v)
	}
}
