package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hidori/go-contexttime"
)

func main() {
	production()
	test()
	production()
	test()
}

func production() {
	ctx := context.Background()
	ctx = contexttime.SetDefaultTime(ctx)

	fmt.Println(contexttime.GetTime(ctx).Now())
}

var _ = contexttime.Time((*MockTime)(nil))

type MockTime struct{}

func (t *MockTime) Now() time.Time {
	return time.Date(2025, 2, 21, 12, 34, 56, 789, time.UTC)
}

func test() {
	ctx := context.Background()
	ctx = contexttime.SetTime(ctx, &MockTime{})

	fmt.Println(contexttime.GetTime(ctx).Now())
}
