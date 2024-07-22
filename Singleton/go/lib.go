package lib

import (
	"sync"
)

// 私有, 防止外部创建实例
type singleton struct {
}

// 协程安全
var once = sync.Once{}
var instance *singleton

func NewSingleton() *singleton {
	// 只执行一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
