package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_set_parse(t *testing.T) {
	tests := []struct {
		line      string
		wantGreen int32
		wantBlue  int32
		wantRed   int32
	}{
		{"6 green, 3 blue", 6, 3, 0},
		{"1 blue, 7 red, 9 green", 9, 1, 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			s := parseSet(tt.line)
			assert.Equal(t, s.blue(), tt.wantBlue)
			assert.Equal(t, s.red(), tt.wantRed)
			assert.Equal(t, s.green(), tt.wantGreen)
		})
	}
}

func Test_set_possibleFor(t *testing.T) {
	tests := []struct {
		s     set
		other set
		want  bool
	}{
		{set{1, 1, 1}, set{2, 2, 2}, true},
		{set{2, 2, 2}, set{1, 1, 1}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.s.possibleFor(tt.other); got != tt.want {
				t.Errorf("set.possibleFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processLine(t *testing.T) {
	tests := []struct {
		line   string
		config set
		want   bool
		wantID int
	}{
		{
			"Game 1: 6 green, 3 blue; 3 red, 1 green; 4 green, 3 red, 5 blue",
			set{100, 100, 100},
			true,
			1,
		},
		{
			"Game 1: 1 green, 1 blue; 1 green, 1 blue",
			set{1, 1, 1},
			false,
			1,
		},
		{
			"Game 100: 6 green, 3 blue; 3 red, 1 green; 4 green, 3 red, 5 blue",
			set{1, 1, 1},
			false,
			100,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if _, got := processLine(tt.line, tt.config); got != tt.want {
				t.Errorf("processLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
