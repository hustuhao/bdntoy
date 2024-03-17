package service

import (
	"fmt"
)

var (
	ErrNotLogin = constError("无效登录，请重新登录")
)

type constError string

func (e constError) Error() string {
	return string(e)
}

// ErrUnexpectedStatusCode is returned on unexpected HTTP status codes
type ErrUnexpectedStatusCode int

func (err ErrUnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected status code: %d", err)
}
