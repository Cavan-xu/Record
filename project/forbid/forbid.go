package forbid

import (
	"strings"
	"sync"
)

type Forbidden struct {
	keywordMap map[rune]*List
}

const (
	UniCodeSpace = 32 // 空格
	UniCodeSpot  = 42 // *
)

var (
	once       sync.Once
	_forbidden = &Forbidden{}
)

func NewForbidden() *Forbidden {
	once.Do(func() {
		_forbidden = &Forbidden{
			keywordMap: make(map[rune]*List),
		}
	})

	return _forbidden
}

func (f *Forbidden) Insert(str string) {
	if len(str) == 0 {
		return
	}

	runes := f.StrTidy(str)
	start := runes[0]

	list, ok := f.keywordMap[start]
	if !ok {
		list = NewList()
		f.keywordMap[start] = list
	}

	node := &Node{value: runes[1:]}
	list.Add(node)
}

func (f *Forbidden) ExactMatchSearch(str string) (string, bool) {
	if len(str) == 0 {
		return "", false
	}

	runes := f.StrTidy(str)
	list, ok := f.keywordMap[runes[0]]
	if !ok {
		return str, false
	}

	if !list.ExactMatchSearch(runes[1:]) {
		return str, false
	}

	for i := 0; i < len(runes); i++ {
		runes[i] = rune(UniCodeSpot)
	}

	return string(runes), true
}

type MatchRes struct {
	start int
	end   int
}

func (f *Forbidden) CommonPrefixSearch(str string) (string, bool) {
	if len(str) == 0 {
		return str, false
	}

	var (
		matched = false
		matches []*MatchRes
		runes   = f.StrTidy(str)
	)

	for i := 0; i < len(runes); i++ {
		list, ok := f.keywordMap[runes[i]]
		if !ok {
			continue
		}
		maxMatch := list.CommonPrefixSearch(runes[i+1:])
		if maxMatch >= 0 {
			matches = append(matches, &MatchRes{start: i, end: i + maxMatch})
			matched = true
			i += maxMatch
		}
	}

	for _, match := range matches {
		for i := match.start; i <= match.end; i++ {
			runes[i] = UniCodeSpot
		}
	}

	return string(runes), matched
}

func (f *Forbidden) StrTidy(str string) []rune {
	str = strings.Trim(str, "\t\r\n")
	runes := []rune(str)
	runesCopy := make([]rune, 0, len(runes))

	for _, r := range runes {
		if r == UniCodeSpace { // 32是空格
			continue
		}
		runesCopy = append(runesCopy, r)
	}

	return runesCopy
}
