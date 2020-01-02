package test

import (
	"log"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ints := make(chan int, 2)
	ints <- 1
	ints <- 1
	log.Println(len(ints))
	//ints <- 1
}

func TestThrottle(t *testing.T) {
	log.Println("start")
	tasks := make(chan float64, 100)

	for i := 0.01; i < 0.3; i += 0.01 {
		tasks <- i
	}
	//go func() {
	//	for i := 0.01; i < 0.3; i += 0.01 {
	//		tasks <- i
	//	}
	//}()
	speed := 10
	c := make(chan int, speed)

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				log.Printf("speed: %d, len: %d", speed, len(c))
				//for i := 0; i < speed-len(c); i++ {
				//	log.Printf("add: %d, len: %d", i, len(c))
				//	c <- i
				//}
				for i := 0; i < speed && len(c) < speed; i++ {
					log.Printf("add: %d", i)
					c <- i
				}
			default:
			}
		}
	}()

	go func() {
		t := time.NewTimer(15 * time.Second)
		select {
		case <-t.C:
			for i := 0.31; i < 0.5; i += 0.01 {
				tasks <- i
			}
		}
	}()

	for {
		select {
		case i := <-c:
			log.Printf("%d - %f", i, <-tasks)
		default:
		}
	}
	//for range c {
	//	log.Printf("%f (%d)", <-tasks, len(c))
	//}
}
