package wordfilter

import (
	"sync"

	"github.com/yicixin/wordfilter/ac"
)

var (
	wordfilter     WordFilter
	wordfilterOnce sync.Once
)

func Init() {
	wordfilterOnce.Do(func() {
		wordfilter = ac.NewAcFilter()
	})
}

func NewWordFilter() WordFilter {
	return ac.NewAcFilter()
}
