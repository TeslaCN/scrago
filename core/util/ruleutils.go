package util

import (
	"github.com/TeslaCN/scrago/cmd/scrago/config"
	"log"
	"regexp"
)

func MatchedRule(rule config.Rule, path string) bool {
	for _, pattern := range rule.UrnPattern {
		m, e := regexp.Match(pattern, []byte(path))
		if e != nil {
			log.Fatalln(e)
		}
		if m {
			return true
		}
	}
	return false
}
