package util

import "testing"

func TestHashCode(t *testing.T) {
	t.Log(HashCode("hello, world"))
}

func TestReserve(t *testing.T) {
	t.Log(Reserve(2<<5-1, 25))
}
