package singleton_test

import (
	"testing"

	singleton "github.com/magicJie/go-design-pattern/c_01_singleton"
	"github.com/stretchr/testify/assert"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetInstance(), singleton.GetInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			if singleton.GetLazyInstance() != singleton.GetLazyInstance() {
				b.Error("test fail")
			}
		}
	})
}
