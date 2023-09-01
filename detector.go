package detector

import (
	_ "embed"
	"strings"
)

//go:embed traditional.dict
var dictString string

var traditionalDict []string

func loadDictionary() {
	traditionalDict = strings.Split(dictString, "\n")
}

func init() {
	loadDictionary()
}

func DetectTraditionalChinese(input string) bool {
	for _, word := range traditionalDict {
		if strings.Contains(input, word) {
			return true
		}
	}
	return false
}
