package pipeline

import (
	"log"
	"scrago/core/pipeline"
)

type DiscardPipeline struct {
}

func (d DiscardPipeline) Process(item interface{}, pipelineHolder pipeline.PipelinesHolder) interface{} {
	log.Printf("Discard -> %s\n", item)
	pipelineHolder.Interrupt()
	return nil
}
