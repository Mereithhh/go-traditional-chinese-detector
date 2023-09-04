package detector

import (
	_ "embed"
	"strings"
	"time"
)

//go:embed traditional.dict
var dictString string

var traditionalDict []string

func loadDictionaryFromEmbedString() {
	traditionalDict = strings.Split(dictString, "\n")
}

func init() {
	loadDictionaryFromEmbedString()
}

type Detector struct {
	dict         []string
	mode         int // 0 内置字典,  1 自定义字典
	loadInterval int // 单位秒，定时重新加载字典，只对自定义字典有效
	getterFn     GetDictFn
}

type GetDictFn func() ([]string, error)

// 新建探测器，使用内置默认字典
func NewDetector() *(Detector) {
	return &Detector{
		dict:         traditionalDict,
		mode:         0,
		loadInterval: 0,
		getterFn:     nil,
	}
}

// 新建探测器，使用自定义字典获取函数
func NewCustomDetector(getDictFn GetDictFn, fetchInterval int) (*Detector, error) {
	dict, err := getDictFn()
	if err != nil {
		return nil, err
	}

	detector := &Detector{
		dict:         dict,
		mode:         1,
		loadInterval: fetchInterval,
		getterFn:     getDictFn,
	}
	if fetchInterval > 0 {
		go func() {
			for {
				detector.ReloadDict()
				time.Sleep(time.Duration(fetchInterval) * time.Second)
			}
		}()
	}
	return detector, nil
}

func (d *Detector) ReloadDict() {
	if d.mode == 0 {
		return
	}
	dict, err := d.getterFn()
	if err != nil {
		return
	}
	d.dict = dict
}

func (d *Detector) Detect(input string) bool {
	return DetectWithDict(input, d.dict)
}

func DetectWithDict(input string, dict []string) bool {
	for _, word := range dict {
		if strings.Contains(input, word) {
			return true
		}
	}
	return false
}

func DetectTraditionalChinese(input string) bool {
	return DetectWithDict(input, traditionalDict)
}
