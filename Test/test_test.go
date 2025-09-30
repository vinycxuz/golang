package Test

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"ana", true},
		{"ovo", true},
		{"radar", true},
		{"reviver", true},
		{"casa", false},
		{"carro", false},
		{"palindrómo", false},
		{"notebook", false},
		{"amor", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func testIsNotPalindrome(t *testing.T) {
	if IsPalindrome("casa") {
		t.Error("casa is not palindrome")
	}
}
