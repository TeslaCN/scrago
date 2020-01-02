package setting

import (
	custom "github.com/TeslaCN/scrago/custom/pipeline"
	"reflect"
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
