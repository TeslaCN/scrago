package rest

import (
	"encoding/json"
	"github.com/TeslaCN/scrago/cmd/scrago/config"
	"log"
	"net/http"
)

func StartRestServer() {
	http.HandleFunc("/config/worker", func(writer http.ResponseWriter, request *http.Request) {
		workerConfig := config.GetWorkerConfig()
		bytes, e := json.Marshal(workerConfig)
		if e != nil {
			log.Println(e)
		}
		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/config/works", func(writer http.ResponseWriter, request *http.Request) {
		configs := config.GetWorkConfig()
		bytes, e := json.Marshal(configs)
		if e != nil {
			log.Println(e)
		}
		_, _ = writer.Write(bytes)
		writer.WriteHeader(http.StatusOK)
	})
	_ = http.ListenAndServe(":6060", nil)
}
