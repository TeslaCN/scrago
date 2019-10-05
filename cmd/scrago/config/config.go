package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type WorkerConfig struct {
	WorkConfig []string          `json:"work_config"`
	Custom     map[string]string `json:"custom"`
}

var (
	configPath = flag.String("c", "", "-c /path/to/config.json")

	workerConfig = &WorkerConfig{}
	workConfigs  []*WorkConfig
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

func GetWorkConfigs() []WorkConfig {
	var works []WorkConfig
	for _, e := range workConfigs {
		works = append(works, *e)
	}
	return works
}

func loadWorkConfig() {
	for _, path := range workerConfig.WorkConfig {
		file, e := os.Open(path)
		if e != nil {
			log.Println("Work Config file not found: " + path)
			continue
		}
		workConfig := &WorkConfig{}
		decoder := json.NewDecoder(file)
		if e := decoder.Decode(workConfig); e != nil {
			log.Println(e)
		}
		log.Println(workConfig)
		workConfigs = append(workConfigs, workConfig)
		_ = file.Close()
	}
}
