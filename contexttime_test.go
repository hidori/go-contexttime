package contexttime

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestNewDefaultTime(t *testing.T) {
	tests := []struct {
		name string
		want *DefaultTime
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultTime_Now(t *testing.T) {
	tests := []struct {
		name string
		tr   *DefaultTime
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &DefaultTime{}
			if got := tr.Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultTime.Now() = %v, want %v", got, tt.want)
			}
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
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetTime(tt.args.ctx, tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTime() = %v, want %v", got, tt.want)
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTime(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTime() = %v, want %v", got, tt.want)
			}
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
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDefaultTime(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDefaultTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
