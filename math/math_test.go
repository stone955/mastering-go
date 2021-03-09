package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative + negative",
			args: args{-1, -1},
			want: -2,
		},
		{
			name: "negative + positive",
			args: args{-1, 1},
			want: 0,
		},
		{
			name: "positive + positive",
			args: args{1, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMulti(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative * negative",
			args: args{-2, -2},
			want: 4,
		},
		{
			name: "negative * positive",
			args: args{-2, 2},
			want: -4,
		},
		{
			name: "positive * positive",
			args: args{2, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multi(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddAssert(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative + negative",
			args: args{-1, -1},
			want: -2,
		},
		{
			name: "negative + positive",
			args: args{-1, 1},
			want: 0,
		},
		{
			name: "positive + positive",
			args: args{1, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Add(tt.args.x, tt.args.y), tt.name)
		})
	}
}

func TestMultiAssert(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "negative * negative",
			args: args{-2, -2},
			want: 4,
		},
		{
			name: "negative * positive",
			args: args{-2, 2},
			want: -4,
		},
		{
			name: "positive * positive",
			args: args{2, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Multi(tt.args.x, tt.args.y), tt.name)
		})
	}
}
