package lib

import (
	"context"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func HostnameOrLocalHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "localhost"
	}
	return hostname
}

// RequestIdKey is the key for the request ID in the context.
var requestIdKey = RequestIdKey

func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(&requestIdKey).(string)
}

func RequestIDWithContext(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, &requestIdKey, value)
}
