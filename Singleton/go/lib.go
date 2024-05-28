package lib

import (
	"sync"
)

// 私有, 防止外部创建实例
type singleton struct {
}

// 协程安全
var mutex = &sync.Mutex{}
var instance *singleton

func NewSingleton() *singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		instance = &singleton{}
	}
	return instance
}
