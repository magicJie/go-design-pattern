package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

// Clone 使用序列化和反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyWord Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyWord)
	return &newKeyWord
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	// 浅拷贝，拷贝了地址
	for k, v := range words {
		newKeywords[k] = v
	}

	// 替换需要更新的字段
	for _, word := range updateWords {
		newKeywords[word.word] = word.Clone()
	}
	return newKeywords
}
