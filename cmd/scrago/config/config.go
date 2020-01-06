package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type WorkerConfig struct {
	WorkConfigPath string `json:"work_config_path"`
	Deduplication  struct {
		Name       string
		Parameters map[string]string
	} `json:"deduplication"`
	Pool struct {
		Name       string
		Parameters map[string]string
	} `json:"pool"`
	Custom map[string]string `json:"custom"`
}

var (
	configPath = flag.String("c", "", "-c /path/to/config.json")

	workerConfig = &WorkerConfig{}
	workConfig   *WorkConfig
)

func init() {

}

func load() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if configPath == nil || *configPath == "" {
		log.Fatalln("specify config path by -c /path/to/config.json")
	}
	file, e := os.Open(*configPath)
	if e != nil {
		log.Fatalln("Config file not found: " + *configPath)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if e := decoder.Decode(workerConfig); e != nil {
		log.Println(e)
	}
	log.Println(*workerConfig)
	loadWorkConfig()
}

func InitConfig() {
	load()
}

func GetWorkerConfig() WorkerConfig {
	return *workerConfig
}

func GetWorkConfig() WorkConfig {
	return *workConfig
}

func loadWorkConfig() {
	file, e := os.Open(workerConfig.WorkConfigPath)
	if e != nil {
		log.Fatalln("Work Config file not found: " + workerConfig.WorkConfigPath)
	}
	wc := &WorkConfig{}
	decoder := json.NewDecoder(file)
	if e := decoder.Decode(wc); e != nil {
		log.Println(e)
	}
	log.Println(wc)
	workConfig = wc
	_ = file.Close()
}
