package util

import (
	"github.com/TeslaCN/scrago/cmd/scrago/config"
	"github.com/TeslaCN/scrago/core/pipeline"
	"github.com/TeslaCN/scrago/core/setting"
	"log"
	"reflect"
)

func GetPipelineByName(name string) (pipeline.Pipeline, bool) {
	i, ok := setting.GetPipelineType(name)
	if !ok {
		log.Fatalf("Pipeline [%s] not found\n", name)
	}
	value := reflect.New(i.Elem())
	p := value.Interface().(pipeline.Pipeline)
	return p, ok
}

func NewPipelineHolder(data interface{}, pipelineConfigs []config.PipelineConfig) pipeline.PipelinesHolder {
	var p []pipeline.Pipeline
	for _, pipelineConfig := range pipelineConfigs {
		pipelineInstance, ok := GetPipelineByName(pipelineConfig.Name)
		if !ok {
			log.Fatalf("Get pipeline [%s] Failed\n", pipelineConfig.Name)
		}
		properties := pipelineConfig.Properties
		value := reflect.ValueOf(pipelineInstance).Elem()
		for k, v := range properties {
			value.FieldByName(k).SetString(v)
		}
		p = append(p, pipelineInstance)
	}
	pipelineHolder := &pipeline.DefaultPipelinesHolder{
		Pipelines:       p,
		CurrentPipeline: 0,
		Data:            &data,
	}
	return pipelineHolder
}
