package singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton != nil {
		return lazySingleton
	}
	once.Do(func() {
		lazySingleton = &Singleton{a: "test"}
	})
	return lazySingleton
}
