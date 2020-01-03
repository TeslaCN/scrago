package setting

import "reflect"

var (
	pipelineMapping     = make(map[string]reflect.Type)
	deduplicationOffset = 25
)

func GetBloomFilterSize() uint64 {
	return 1 << deduplicationOffset
}

func SetDeduplicationOffset(offset int32) {
	deduplicationOffset = int(offset)
}

func GetDeduplicationOffset() int32 {
	return int32(deduplicationOffset)
}

func AddPipelineMapping(name string, p reflect.Type) {
	pipelineMapping[name] = p
}

func GetPipelineMapping() map[string]reflect.Type {
	return pipelineMapping
}
