package ref

import (
	"fmt"
	"github.com/TeslaCN/scrago/core/reflection"
	"github.com/TeslaCN/scrago/custom/item"
	"net/http"
	"reflect"
	"testing"
)

type Pipeline interface {
	handle(request http.Request, response http.Response, next Pipeline)
}

type ElasticsearchPipeline struct {
}

func (*ElasticsearchPipeline) handle(request http.Request, response http.Response, next Pipeline) {

}

type DiscardPipeline struct {
}

func (*DiscardPipeline) handle(request http.Request, response http.Response, next Pipeline) {

}

func TestReflect(t *testing.T) {

}

func TestParseStruct(t *testing.T) {
	parseStruct := reflection.ParseStruct(item.VideoInformationTemp{})
	fmt.Println(parseStruct)
}

type TestStruct struct {
	Name  string
	Value string
}

func BenchmarkDirectValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := TestStruct{}
		s.Name = "Sia"
		s.Value = "hello, world"
		s.Name = "Sia"
		s.Value = "hello, world"
		s.Name = "Sia"
		s.Value = "hello, world"
		s.Name = "Sia"
		s.Value = "hello, world"
		s.Name = "Sia"
		s.Value = "hello, world"
		s.Name = "Sia"
		s.Value = "hello, world"
	}
}

func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := TestStruct{}
		t := reflect.ValueOf(&s).Elem()
		t.FieldByName("Name").SetString("Sia")
		t.FieldByName("Value").SetString("hello, world")
		t.FieldByName("Name").SetString("Sia")
		t.FieldByName("Value").SetString("hello, world")
		t.FieldByName("Name").SetString("Sia")
		t.FieldByName("Value").SetString("hello, world")
		t.FieldByName("Name").SetString("Sia")
		t.FieldByName("Value").SetString("hello, world")
		t.FieldByName("Name").SetString("Sia")
		t.FieldByName("Value").SetString("hello, world")
		t.FieldByName("Name").SetString("Sia")
		t.FieldByName("Value").SetString("hello, world")
	}
}

func BenchmarkN(b *testing.B) {
	n := 0
	for i := 0; i < b.N; i++ {
		n++
	}
}
