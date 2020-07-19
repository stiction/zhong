package zhong

import "testing"

func TestCheckMobile(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", false},
		{"hello", false},
		{"13800 38000", false},
		{"13800a38000", false},
		{"11800138000", false},
		{"12800138000", false},
		{"23800138000", false},
		{"1380013800", false},
		{"138001380001", false},
		{"13800138000 ", false},
		{" 13800138000", false},
		{"\n13800138000", false},
		{"13800138000\n", false},
		{"13800138000", true},
		{"14800138000", true},
		{"15800138000", true},
		{"16800138000", true},
		{"17800138000", true},
		{"18800138000", true},
		{"19800138000", true},
	}
	for _, test := range tests {
		if got := CheckMobile(test.input); got != test.want {
			t.Errorf("CheckMobile(%q) = %v", test.input, got)
		}
	}
}
