package lib

// 私有, 防止外部创建实例
type singleton struct {
}

var instance *singleton

func NewSingleton() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
