package config

import "testing"

func TestMap(t *testing.T) {

	strings := make(map[string]string)
	strings["es.hosts"] = "http://es.0:49204"

	InitConfig()
	t.Log(GetWorkerConfig().Custom["elasticsearch.hosts"])
	t.Log(strings["es.hosts"])
	t.Log(strings["es.username"] == "")
}
