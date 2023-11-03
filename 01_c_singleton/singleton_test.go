package singleton_test

import (
	"testing"

	singleton "github.com/magicJie/go-design-pattern/c_01_singleton"
	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.True(t, singleton.GetInstance() == singleton.GetInstance())
	assert.False(t, singleton.GetInstance() == singleton.GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Error("test fail")
			}
		}
	})
}
