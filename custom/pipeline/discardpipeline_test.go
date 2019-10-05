package pipeline

import (
	"fmt"
	"reflect"
	"scrago/core/pipeline"
	"testing"
)

func BenchmarkCreateDirectly(b *testing.B) {
	var a []*DiscardPipeline
	for i := 0; i < b.N; i++ {
		a = append(a, &DiscardPipeline{})
	}
}

func BenchmarkReflectCreate(b *testing.B) {
	t := reflect.TypeOf(&DiscardPipeline{})
	var a []*DiscardPipeline
	for i := 0; i < b.N; i++ {
		value := reflect.New(t.Elem())
		a = append(a, value.Interface().(*DiscardPipeline))
	}
}

func TestPrintBytes(t *testing.T) {
	s := "hello, world"

	b := []byte(s)
	fmt.Println(b)
	fmt.Printf("%s\n", b)
	fmt.Println(s)
}

func TestCast(t *testing.T) {
	var p pipeline.Pipeline
	pipelineType := reflect.TypeOf(&DiscardPipeline{}).Elem()
	p = reflect.New(pipelineType).Interface().(pipeline.Pipeline)
	fmt.Println(p)
}
