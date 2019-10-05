package pipeline

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"scrago/core/net"
	"scrago/core/pipeline"
	"scrago/core/reflection"
	"scrago/custom/item"
	"scrago/custom/pipeline/util"
)

type JavInfoPipeline struct {
}

func (j JavInfoPipeline) Process(i interface{}, pipelineHolder pipeline.PipelinesHolder) interface{} {
	httpResponse, ok := (*i.(*interface{})).(*net.HttpResponse)
	if !ok {
	}
	document, e := goquery.NewDocumentFromReader(bytes.NewReader(httpResponse.Body))
	if e != nil {
	}
	detail := item.VideoInformationTemp{}
	metaFields := reflection.ParseStruct(detail)
	m := make(map[string]interface{})
	for field, meta := range metaFields {
		selection := document.Find(meta.CssSelector)
		var value interface{}
		if selection.Length() > 1 {
			value = selection.Map(func(i int, s *goquery.Selection) string {
				if meta.Attr != "" {
					val, _ := s.Attr(meta.Attr)
					return val
				}
				return s.Text()
			})
		} else {
			if meta.Attr != "" {
				value, _ = selection.Attr(meta.Attr)
			} else {
				value = selection.Text()
			}
		}
		m[field] = value
	}

	if scripts, ok := m["torrent_scripts"].([]string); ok {
		var torrents []string
		for _, s := range scripts {
			if s == "" {
				continue
			}
			torrents = append(torrents, util.ParseJs(s))
		}
		m["torrents"] = torrents
		delete(m, "torrent_scripts")
	}
	bytes, _ := json.Marshal(m)
	pipelineHolder.SetData(bytes)
	pipelineHolder.Next()
	return bytes
}

type JavBookTorrentDecodePipeline struct {
}

// magnet:?xt=urn:btih:0caee55840ae0e4dd7bdf97ac6804b6bdcd12cd3
// 0caee55840ae0e4dd7bdf97ac6804b6bdcd12cd3
// 00111010011011010110101101101111011011110011111100111111010000100011111000111010011010110110111100111010011011110011111001101110011011100100000101101100011011100111000001000011010000010110101101101101010000000100001000111010001111100110110001000000011011000110111001101101011011100011101100111100011011010110111000111101
func (JavBookTorrentDecodePipeline) Process(item interface{}, pipelineHolder pipeline.PipelinesHolder) interface{} {

	return nil
}
