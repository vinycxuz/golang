package Test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestSum(t *testing.T) {
	got := Sum(4, 4)
	if got != 8 {
		t.Errorf("Sum(4, 4) = %d; esperado 8", got)
	}
}

func TestNewPerson(t *testing.T) {
	t.Run("Idade positiva", func(t *testing.T) {
		age := 25
		got, err := NewPerson(age)
		if err != nil {
			t.Errorf("NewPerson(%d) retornou um erro inesperado: %v", age, err)
		}
		if got == nil {
			t.Fatal("NewPerson retornou nil")
		}
		if got.age != age {
			t.Errorf("NewPerson(%d) = %v; esperado idade %d", age, got.age, age)
		}
	})

	t.Run("Idade negativa", func(t *testing.T) {
		age := -5
		expectedErr := "idade não pode ser negativa"
		got, err := NewPerson(age)
		if err == nil {
			t.Error("NewPerson não retornou erro para idade negativa")
		}
		if err.Error() != expectedErr {
			t.Errorf("NewPerson(%d) retornou erro %q; esperado %q", age, err.Error(), expectedErr)
		}
		if got != nil {
			t.Errorf("NewPerson(%d) = %v; esperado nil", age, got)
		}
	})
}

func TestTableNewPerson(t *testing.T) {
	testSuite := []struct {
		name       string
		input      int
		expected   int
		errMessage string
	}{
		{
			name:       "Idade positiva",
			input:      30,
			expected:   30,
			errMessage: "",
		},
		{
			name:       "Idade negativa",
			input:      -10,
			expected:   0,
			errMessage: "idade não pode ser negativa",
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPerson(tt.input)
			if tt.errMessage != "" {
				if err == nil {
					t.Fatalf("NewPerson(%d) retornou erro inesperado", tt.input)
				}
				if err.Error() != tt.errMessage {
					t.Errorf("NewPerson(%d) retornou erro %q; esperado %q", tt.input, err.Error(), tt.errMessage)
				}
				if got != nil {
					t.Errorf("NewPerson(%d) = %v; esperado nil", tt.input, got)
				}
				return
			}
			if err != nil {
				t.Errorf("NewPerson(%d) retornou erro inesperado: %v", tt.input, err)
			}
			if got == nil {
				t.Fatal("NewPerson retornou nil")
			}
			if got.age != tt.expected {
				t.Errorf("NewPerson(%d) = %v; esperado idade %d", tt.input, got.age, tt.expected)
			}
		})
	}
}

func TestQuickFunc(t *testing.T) {
	input := "hello"
	got := QuickFunc(input)
	if !strings.Contains(got, input) {
		t.Errorf("QuickFunc(%q) = %q; esperado conter %q", input, got, input)
	}
}

func TestSlowFunc(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	input := "hi"
	got := SlowFunc(input)
	if !strings.Contains(got, input) {
		t.Errorf("SlowFunc(%q) = %q; esperado conter %q", input, got, input)
	}
}

func TestNewPerson_WithTenAsAge(t *testing.T) {
	got, err := NewPerson(10)
	assert.NotNil(t, got)
	assert.Nil(t, err)
	assert.Equal(t, 10, got.age)
}
