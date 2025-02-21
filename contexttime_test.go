package contexttime

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultTime(t *testing.T) {
	tests := []struct {
		name string
		want *DefaultTime
	}{
		{
			name: "success: returns DefaultTime",
			want: &DefaultTime{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDefaultTime()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDefaultTime_Now(t *testing.T) {
	tests := []struct {
		name string
		tr   *DefaultTime
		want time.Duration
	}{
		{
			name: "success: returns current time",
			tr:   &DefaultTime{},
			want: 1 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &DefaultTime{}
			got := tr.Now()
			assert.LessOrEqual(t, time.Since(got), tt.want)
		})
	}
}

func TestSetTime(t *testing.T) {
	type args struct {
		ctx  context.Context
		time Time
	}
	tests := []struct {
		name string
		args args
		want Time
	}{
		{
			name: "success: attach Time to context.Context",
			args: args{
				ctx:  context.Background(),
				time: &DefaultTime{},
			},
			want: &DefaultTime{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetTime(tt.args.ctx, tt.args.time).Value(KeyContextTime)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTime(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Time
	}{
		{
			name: "success: returns Time",
			args: args{
				ctx: context.WithValue(context.Background(), KeyContextTime, &DefaultTime{}),
			},
			want: &DefaultTime{},
		},
		{
			name: "success: returns nil",
			args: args{
				ctx: context.Background(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTime(tt.args.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSetDefaultTime(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Time
	}{
		{
			name: "success: returns Time",
			args: args{
				ctx: context.Background(),
			},
			want: &DefaultTime{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SetDefaultTime(tt.args.ctx).Value(KeyContextTime)
			assert.Equal(t, tt.want, got)
		})
	}
}
