package setting

import "reflect"

var (
	pipelineMapping      = make(map[string]reflect.Type)
	deduplicationMapping = make(map[string]reflect.Type)
	poolMapping          = make(map[string]reflect.Type)
	deduplicationOffset  = 25
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

func AddPipelineType(name string, p reflect.Type) {
	pipelineMapping[name] = p
}

func GetPipelineType(name string) (reflect.Type, bool) {
	t, ok := pipelineMapping[name]
	return t, ok
}

func AddDeduplicationType(name string, d reflect.Type) {
	deduplicationMapping[name] = d
}
func GetDeduplicateType(name string) (reflect.Type, bool) {
	t, ok := deduplicationMapping[name]
	return t, ok
}

func AddPoolType(name string, p reflect.Type) {
	poolMapping[name] = p
}

func GetPoolType(name string) (reflect.Type, bool) {
	t, ok := poolMapping[name]
	return t, ok
}
