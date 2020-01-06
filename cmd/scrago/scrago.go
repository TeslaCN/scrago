package scrago

import (
	"context"
	"github.com/TeslaCN/scrago/cmd/scrago/config"
	"github.com/TeslaCN/scrago/cmd/scrago/worker"
	"github.com/TeslaCN/scrago/core/adapter/rest"
	"log"
	"os"
	"os/signal"
	"sync"
)

var (
	wg = new(sync.WaitGroup)
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Start() {
	c := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "wg", wg))
	signal.Notify(c, os.Interrupt, os.Kill)

	config.InitConfig()

	worker.StartWorker(ctx)

	go rest.StartRestServer()

	s := <-c
	log.Printf("System Signal: %s", s)
	cancel()
	wg.Wait()
}
