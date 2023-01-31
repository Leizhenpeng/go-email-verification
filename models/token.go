package models

import "time"

// Token token返回结构体
type Token struct {
	Token   string    `json:"token"`   // token
	Expires time.Time `json:"expires"` // 过期时间
}
