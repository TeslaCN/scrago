package pipeline

type Pipeline interface {
	Process(item interface{}, pipelineHolder PipelinesHolder) interface{}
}

type PipelinesHolder interface {
	Next()
	Interrupt()
	Ended() (ended bool)
	GetData() interface{}
	SetData(interface{})
}

type DefaultPipelinesHolder struct {
	Pipelines       []Pipeline
	CurrentPipeline int
	Data            interface{}
}

func (h *DefaultPipelinesHolder) SetData(data interface{}) {
	h.Data = data
}

func (h *DefaultPipelinesHolder) GetData() interface{} {
	return h.Data
}

func (h *DefaultPipelinesHolder) Ended() bool {
	return h.CurrentPipeline == len(h.Pipelines)
}

func (h *DefaultPipelinesHolder) Next() {
	if h.CurrentPipeline < len(h.Pipelines) {
		h.CurrentPipeline++
		processed := h.Pipelines[h.CurrentPipeline-1].Process(h.GetData(), h)
		h.Data = processed
	}
}

func (h *DefaultPipelinesHolder) Interrupt() {
	h.CurrentPipeline = len(h.Pipelines)
}
