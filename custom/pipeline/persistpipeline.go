package pipeline

import (
	"bufio"
	"github.com/TeslaCN/scrago/core/net"
	"github.com/TeslaCN/scrago/core/pipeline"
	"log"
	"os"
	"strings"
)

type PersistLocalFilePipeline struct {
	BasePath string
}

func (p PersistLocalFilePipeline) Process(item interface{}, pipelineHolder pipeline.PipelinesHolder) interface{} {
	//httpResponse, ok := (*item.(*interface{})).(map[string]interface{})
	httpResponse, ok := (*item.(*interface{})).(*net.HttpResponse)
	split := strings.Split(httpResponse.Uri, "/")
	var fileName string
	if len(split) > 0 {
		fileName = split[len(split)-1]
	} else {

	}
	filePath := p.BasePath + fileName
	if !ok {
		pipelineHolder.Interrupt()
		return nil
	}

	file, e := os.Create(filePath)
	if e != nil {
		log.Printf("Error Creating [%s] %v\n", fileName, e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	_, _ = writer.Write(httpResponse.Body)
	log.Printf("Persist Image [%s] -> [%s]\n", httpResponse.Uri, filePath)
	pipelineHolder.Next()
	return httpResponse
}

type PersistElasticsearchPipeline struct {
}
