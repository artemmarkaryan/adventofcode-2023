package main

import (
	"testing"
)

func Test_firstNumber(t *testing.T) {
	tests := []struct {
		tokens []token
		line   string
		want   int32
	}{
		{straigntTokens, "one", 1},
		{straigntTokens, "oneasdfasd", 1},
		{straigntTokens, "onetwo", 1},
		{reversedTokens, "eno", 1},
		{reversedTokens, "asdfsdfeno", 1},
		{reversedTokens, "owteno", 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := firstNumber(tt.line, tt.tokens); got != tt.want {
				t.Errorf("firstNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLine(t *testing.T) {
	tests := []struct {
		line string
		want int32
	}{
		{"onetwo", 12},
		{"onetwothree", 13},
		{"1twothree", 13},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := calculateLine(tt.line); got != tt.want {
				t.Errorf("calculateLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
