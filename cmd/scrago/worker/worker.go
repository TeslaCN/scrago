package worker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/TeslaCN/scrago/cmd/scrago/config"
	"github.com/TeslaCN/scrago/core/net"
	"github.com/TeslaCN/scrago/core/task"
	"github.com/TeslaCN/scrago/core/util"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

type Worker interface {
	Start(ctx context.Context)
}

type DefaultWorker struct {
	WorkerName  string
	Work        config.Work
	ctx         context.Context
	pool        task.Pool
	deduplicate task.Deduplicate
	requester   net.HttpRequester
}

func (w *DefaultWorker) Start(ctx context.Context) {
	w.ctx = ctx
	waitGroup := w.ctx.Value("wg").(*sync.WaitGroup)
	waitGroup.Add(1)

	// add seeds into queue
	for _, seed := range w.Work.Seeds {
		parsed, e := url.Parse(seed)
		if e != nil || parsed == nil {
			log.Printf("[%s] Seed [%s] %s\n", w.WorkerName, seed, e)
			continue
		}
		w.pool.Offer(task.Task{Url: *parsed})
	}

	goroutines := runtime.NumCPU()
	// goroutines := 1
	for i := 0; i < goroutines; i++ {
		go w.doWork(i)
	}
}

func (w *DefaultWorker) doWork(num int) {
	waitGroup := w.ctx.Value("wg").(*sync.WaitGroup)
	defer waitGroup.Done()

	workerName := fmt.Sprintf("%s - %d", w.WorkerName, num)

	defer log.Printf("[%s] Stopped\n", workerName)

	log.Printf("[%s] Started\n", workerName)

WORKING:
	for {
		select {
		case <-w.ctx.Done():
			break WORKING
		default:
			var httpResponse *net.HttpResponse

			// fetch task
			fetchedTask, e := w.pool.Fetch()
			if e != nil {

			}

			// deduplicate
			fetchedTimes := w.deduplicate.De(fetchedTask.Url)
			if fetchedTimes > 0 {
				continue
			}

			log.Printf("[%s] Fetching [%s]\n", workerName, fetchedTask.Url.String())

			// http request
			switch fetchedTask.Method {
			case http.MethodGet:
				httpResponse, e = w.requester.Get(fetchedTask.Url.String(), nil)
			case http.MethodPost:
			default:
				httpResponse, e = w.requester.Get(fetchedTask.Url.String(), nil)
			}
			if e != nil {
				log.Printf("Request [%s] failed", e.Error())
				w.deduplicate.Remove(fetchedTask.Url)
				go w.pool.Offer(*fetchedTask)
				continue
			}

			contentType := httpResponse.Headers.Get("Content-Type")

			var document *goquery.Document
			if isHtml(contentType) {
				document, e = goquery.NewDocumentFromReader(bytes.NewReader(httpResponse.Body))
				if e != nil {
					log.Println(e)
				}
			}

			// put raw into pipeline & create new tasks
			anyMatched := false
			for _, rule := range w.Work.Rules {

				matched := false
				for _, pattern := range rule.UrnPattern {
					m, e := regexp.Match(pattern, []byte(fetchedTask.Url.Path))
					if e != nil {
						log.Fatalln(e)
					}
					if m {
						matched = true
						break
					}
				}

				if !matched {
					continue
				}
				anyMatched = true

				var (
					urls []*url.URL
					f    = rule.Follow
				)

				if isHtml(contentType) {
					// parse all links
					document.Find("a").Each(func(i int, s *goquery.Selection) {
						link, exists := s.Attr("href")
						if !exists {
							return
						}
						u, e := url.Parse(link)
						if e != nil {
							return
						}

						// Deny Domain
						for _, d := range f.DenyDomains {
							if d == u.Host {
								return
							}
						}

						// Allow Domain
						allowDomain := false
						for _, d := range f.AllowDomains {
							if d == u.Host {
								allowDomain = true
							}
						}
						if !allowDomain {
							return
						}

						// Deny Pattern
						for _, r := range f.DenyRules {
							switch r.Type {
							case "regexp":
								matched, e := regexp.Match(r.Value, []byte(u.RequestURI()))
								if e != nil {
									log.Fatalln(e)
								}
								if matched {
									return
								}
							}
						}

						// Allow Pattern
						allowPattern := false
						for _, r := range f.AllowRules {
							switch r.Type {
							case "regexp":
								matched, e := regexp.Match(r.Value, []byte(u.RequestURI()))
								if e != nil {
									log.Fatalln(e)
								}
								if matched {
									allowPattern = true
								}
							}
						}
						if !allowPattern {
							return
						}
						if w.deduplicate.Exist(*u) > 0 {
							return
						}

						urls = append(urls, u)
					})
					// Create & Offer Tasks
					go func() {
						for _, todoUrl := range urls {
							w.pool.Offer(task.Task{
								Method: http.MethodGet,
								Url:    *todoUrl,
							})
						}
					}()
				}

				// Put into Pipeline
				go util.NewPipelineHolder(httpResponse, rule.Pipelines).Next()
			}

			if !anyMatched {
				log.Printf("None Matched [%s]\n", fetchedTask.Url.String())
			}

		}
	}

}

func isHtml(contentType string) bool {
	return strings.Contains(contentType, "text/html")
}

func NewWorker(work config.Work) Worker {
	log.Printf("New Worker [%s]\n", work.Name)
	return &DefaultWorker{
		WorkerName:  work.Name,
		Work:        work,
		pool:        task.NewPool(),
		deduplicate: task.NewDeduplicate(),
		requester:   &net.DefaultHttpRequester{Client: &http.Client{}},
	}
}

func StartWorker(ctx context.Context) {
	log.Println("Starting Workers.")
	workConfigs := config.GetWorkConfigs()
	for _, workConfig := range workConfigs {
		for _, work := range workConfig.Works {
			NewWorker(*work).Start(ctx)
		}
	}
}
