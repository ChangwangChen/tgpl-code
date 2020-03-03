package word_test

import (
	"testing"
	"tgpl-code/chapter11/word"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct{
		input string
		want bool
	} {
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kkakk", true},
		{"Evil I did dwell; lewd did I live.", true},
	}

	for _, test := range tests {
		if got := word.IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.IsPalindrome("A man, a plan, a canal: Panama")
	}
}