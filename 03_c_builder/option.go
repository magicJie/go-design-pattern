package builder

import "fmt"

type ResourcePoolConfigOption struct {
	maxTotal int
	maxIdle  int
	minIdle  int
}

type ResourcePoolConfigOptionFunc func(option *ResourcePoolConfigOption)

func NewResoucePoolConfig(name string, opts ...ResourcePoolConfigOptionFunc) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, fmt.Errorf("name cant not by empty")
	}

	option := &ResourcePoolConfigOption{
		maxTotal: 10,
		maxIdle:  9,
		minIdle:  1,
	}

	for _, opt := range opts {
		opt(option)
	}

	if option.maxTotal < 0 || option.maxIdle < 0 || option.minIdle < 0 {
		return nil, fmt.Errorf("args arr, option: %v", option)
	}

	if option.maxTotal < option.maxIdle || option.minIdle > option.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", option)
	}

	return &ResourcePoolConfig{
		name:     name,
		maxTotal: option.maxTotal,
		maxIdle:  option.maxIdle,
		minIdle:  option.minIdle,
	}, nil
}
