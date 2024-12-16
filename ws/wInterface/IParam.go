package wInterface

import (
	. "github.com/phyer/v5sdkgo/ws/wImpl"
)

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
