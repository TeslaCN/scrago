package concurrent

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	method := func(ctx context.Context, name string) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("%s, exit.\n", name)
				return
			default:
				fmt.Printf("%s, watching...\n", name)
				time.Sleep(2 * time.Second)
			}
		}
	}
	go method(ctx, "1")
	go method(ctx, "2")
	go method(ctx, "3")
	time.Sleep(2500 * time.Millisecond)
	cancel()
	time.Sleep(2 * time.Second)
}

func TestSpeedLimit(t *testing.T) {

	capacity := 5
	c := make(chan int, capacity)

	log.Println("Before => ", cap(c), len(c))
	c <- 1
	log.Println("After => ", cap(c), len(c))

	for i := 0; i < 100; i++ {
		var f func(int)
		f = func(n int) {
			s := <-c
			//if s != n {
			//	c <- s
			//	f(i)
			//}
			log.Printf("%d => %d", s, n)
		}
		go f(i)
	}
	ticker := time.NewTicker(time.Second)
	seq := 0
	for {
		select {
		case <-ticker.C:
			for a := 0; a < capacity-len(c); a++ {
				c <- seq
				seq++
			}
			//case <-ticker.C:
			//	fmt.Println("Now => ", cap(c), len(c))
			//	time.Sleep(500 * time.Millisecond)

			//default:
			//	fmt.Println("Now => ", cap(c), len(c))
			//	time.Sleep(500 * time.Millisecond)
		}
	}
}
