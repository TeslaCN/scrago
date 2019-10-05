package pipeline

import (
	"fmt"
	"scrago/core/pipeline"
)

type Html2StructPipeline struct {
}

func (p *Html2StructPipeline) Process(item interface{}, pipelineHolder pipeline.PipelinesHolder) interface{} {
	var bodyBytes []byte
	b, ok := item.([]byte)
	if ok {
		bodyBytes = b
	}
	if s, ok := item.(string); ok {
		bodyBytes = []byte(s)
	}
	fmt.Println(bodyBytes)
	return nil
}
