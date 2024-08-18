package utils

import (
	"context"
	"time"
)

// RetryFunc This function will be retried multiple times until no errors are reported
func RetryFunc(ctx context.Context, f func() error, retryTimes int, sleepTime time.Duration) {
	for i := 0; i < retryTimes; i++ {
		if err := f(); err != nil {
			time.Sleep(sleepTime)
			continue
		}
		return
	}
}
