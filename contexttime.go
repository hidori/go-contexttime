package contexttime

import (
	"context"
	"time"
)

var _ = Time((*DefaultTime)(nil))

type Time interface {
	Now() time.Time
}

type DefaultTime struct{}

func NewDefaultTime() *DefaultTime {
	return &DefaultTime{}
}

func (t *DefaultTime) Now() time.Time {
	return time.Now()
}

type KeyType string

const KeyContextTime KeyType = "contexttime.Time"

func SetTime(ctx context.Context, time Time) context.Context {
	return context.WithValue(ctx, KeyContextTime, time)
}

func GetTime(ctx context.Context) Time {
	time, ok := ctx.Value(KeyContextTime).(Time)
	if !ok {
		return nil
	}

	return time
}

func SetDefaultTime(ctx context.Context) context.Context {
	return SetTime(ctx, NewDefaultTime())
}
