package test

import (
	"log"
	"testing"
)

func TestMake(t *testing.T) {
	a := make([]int, 3)
	log.Printf("%v, %d, %d", a, len(a), cap(a))

	b := make([]int, 0, 3)
	log.Printf("%v, %d, %d", b, len(b), cap(b))
}
