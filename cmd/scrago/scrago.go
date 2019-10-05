package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"scrago/cmd/scrago/config"
	"scrago/cmd/scrago/worker"
	"sync"
)

var (
	wg = new(sync.WaitGroup)
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	c := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "wg", wg))
	signal.Notify(c, os.Interrupt, os.Kill)

	config.InitConfig()

	worker.StartWorker(ctx)

	s := <-c
	log.Printf("System Signal: %s", s)
	cancel()
	wg.Wait()
}
