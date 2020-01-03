package net

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/PuerkitoBio/goquery"
//	"github.com/TeslaCN/scrago/core/reflection"
//	//"github.com/TeslaCN/scrago/custom/item"
//	//"github.com/TeslaCN/scrago/custom/pipeline/util"
//	"log"
//	"net/http"
//	"net/url"
//	"strings"
//	"testing"
//)
//
//func TestParseUrl(t *testing.T) {
//	u := "https://ss9874.com/content_censored/212678.htm"
//	parsed, e := url.Parse(u)
//	if e != nil {
//		t.Log(e)
//	}
//	t.Log(parsed)
//}
//
//func TestDefaultHttpGet(t *testing.T) {
//	requester := &DefaultHttpRequester{}
//	b, _ := requester.Get("https://ss9874.com/content_censored/212678.htm", nil)
//	fmt.Println(string(b.Body))
//}
//
//func TestGetAndParse(t *testing.T) {
//	requester := &DefaultHttpRequester{}
//	b, _ := requester.Get("https://ss9874.com", nil)
//	fmt.Println(string(b.Body))
//	document, e := goquery.NewDocumentFromReader(bytes.NewReader(b.Body))
//	if e != nil {
//
//	}
//	document.Find(".Po_topic").Each(func(i int, s *goquery.Selection) {
//		code := s.Find("div[class='Po_topic_Date_Serial'] font").Text()
//		title := s.Find("div[class='Po_topic_title'] a b").Text()
//		uri, _ := s.Find("div[class='Po_topicCG'] a").Attr("href")
//		img := s.Find("div[class='Po_topicCG'] a img")
//		smallImage, _ := img.Attr("src")
//		largeImage, _ := img.Attr("onmouseover")
//		largeImage = strings.TrimSuffix(strings.TrimPrefix(largeImage, "showtrail('"), "','',10,10)")
//
//		fmt.Printf("%s %s %s\n", code, title, uri)
//		fmt.Printf("%s %s\n", smallImage, largeImage)
//	})
//}
//
//func TestParsedGet(t *testing.T) {
//	requester := &DefaultHttpRequester{Client: &http.Client{}}
//	b, _ := requester.Get("https://ss9874.com/content_censored/212678.htm", nil)
//	fmt.Println(string(b.Body))
//	document, e := goquery.NewDocumentFromReader(bytes.NewReader(b.Body))
//	if e != nil {
//
//	}
//	// var links []string
//	// document.Find("a").Each(func(i int, selection *goquery.Selection) {
//	// 	val, exists := selection.Attr("href")
//	// 	if !exists {
//	// 		return
//	// 	}
//	// 	links = append(links, val)
//	// })
//	// fmt.Println(links)
//	// fmt.Println(len(links))
//	detail := item.VideoInformationTemp{}
//	metaFields := reflection.ParseStruct(detail)
//	m := make(map[string]interface{})
//	for field, meta := range metaFields {
//		selection := document.Find(meta.CssSelector)
//		var value interface{}
//		if selection.Length() > 1 {
//			value = selection.Map(func(i int, s *goquery.Selection) string {
//				if meta.Attr != "" {
//					val, _ := s.Attr(meta.Attr)
//					return val
//				}
//				return s.Text()
//			})
//		} else {
//			if meta.Attr != "" {
//				value, _ = selection.Attr(meta.Attr)
//			} else {
//				value = selection.Text()
//			}
//		}
//		m[field] = value
//	}
//	scripts := m["torrent_scripts"].([]string)
//	var torrents []string
//	for _, s := range scripts {
//		if s == "" {
//			continue
//		}
//		torrents = append(torrents, util.ParseJs(s))
//	}
//	m["torrents"] = torrents
//	log.Println(m)
//	bytes, _ := json.Marshal(m)
//	log.Println(string(bytes))
//	log.Println()
//}
