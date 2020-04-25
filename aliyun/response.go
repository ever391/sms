package aliyun

import (
	"time"
)

type ResponseMessage struct {
	Mobile      string    `json:"mobile"`
	Message     string    `json:"Message"`
	RequestID   string    `json:"RequestId"`
	BizID       string    `json:"BizId"`
	Code        string    `json:"Code"`
	CreatedTime time.Time `json:"created_time"`
}
