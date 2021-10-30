package db

import (
	"context"
	"time"
)

func GetContextWithDefaultTimeout() (context.Context, context.CancelFunc) {
	return GetContextWithTimeout(5 * time.Second)
}

func GetContextWithTimeout(time time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time)
}
