package main

import (
	"testing"
)

func Test_calculate(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "",
			want:  0,
		}, {
			input: "1",
			want:  11,
		}, {
			input: "11",
			want:  11,
		}, {
			input: "a11",
			want:  11,
		}, {
			input: "1a1",
			want:  11,
		}, {
			input: "11a",
			want:  11,
		}, {
			input: "11a\n11",
			want:  11 + 11,
		}, {
			input: "kfxone67bzb2\n8jjpseven\n236twoknbxlczgd",
			want:  62 + 88 + 26,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := calculate([]byte(tt.input)); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
