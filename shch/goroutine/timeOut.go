package main

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout")

//限定方法执行时间
func RunWithTimeout(handler func(), timeout time.Duration) error {
	done := make(chan bool, 1)
	go func() {
		handler()
		done <- true
		close(done)
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case <-timer.C:
		return ErrTimeout
	case <-done:
		return nil
	}
}
