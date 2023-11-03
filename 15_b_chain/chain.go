package chain

type SensitiveWordFilter interface {
	Filter(content string) bool
}

// 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// 执行过滤
func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

// 广告
type AdSensitiveWordFilter struct{}

// 实现过滤算法
func (f *AdSensitiveWordFilter) Filter(content string) bool {
	return false
}

// 政治敏感
type PoliticalWordFilter struct{}

// 实现过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	return true
}
