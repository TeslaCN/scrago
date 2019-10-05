package item

import (
	"reflect"
	"testing"
)

func TestReflectStruct(t *testing.T) {
	information := VideoInformation{}
	itemType := reflect.TypeOf(information)
	numField := itemType.NumField()
	for i := 0; i < numField; i++ {
		// field := itemType.Field(i)
		// field.Type.
	}
}
