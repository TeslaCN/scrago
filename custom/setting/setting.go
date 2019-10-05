package setting

import (
	"reflect"
	custom "scrago/custom/pipeline"
)

var (
	PipelineMapping = make(map[string]reflect.Type)
)

const DeduplicationOffset = 25
const BloomFilterSize uint64 = 1 << DeduplicationOffset

func init() {
	PipelineMapping["DiscardPipeline"] = reflect.TypeOf(&custom.DiscardPipeline{})
	PipelineMapping["PersistLocalFilePipeline"] = reflect.TypeOf(&custom.PersistLocalFilePipeline{})
	PipelineMapping["JavBookTorrentDecodePipeline"] = reflect.TypeOf(&custom.JavBookTorrentDecodePipeline{})
	PipelineMapping["JavInfoPipeline"] = reflect.TypeOf(&custom.JavInfoPipeline{})
}
