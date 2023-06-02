package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Only run in mac")
	}

	fmt.Println("TestSkip done")
}

func TestHelloRequire(t *testing.T) {
	result := Hello("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")

	fmt.Println("TestHelloRequire done")
}

func TestHello(t *testing.T) {
	result := Hello("Teza")
	if result != "Hello Teza" {
		t.Fail()
	}

	fmt.Println("TestHello done")
}

func TestHelloAssert(t *testing.T) {
	result := Hello("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")

	fmt.Println("TestHelloAssert done")
}

func TestSubTest(t *testing.T) {
	t.Run("EKO", func(t *testing.T) {
		result := Hello("EKO")
		require.Equal(t, "Hello EKO", result, "Return must be 'Hello EKO'")
	})
	t.Run("TEZA", func(t *testing.T) {
		result := Hello("TEZA")
		require.Equal(t, "Hello TEZA", result, "Return must be 'Hello TEZA'")
	})

	fmt.Println("TestSubTest done")
}

func TestHelloTable(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello('TEZA')",
			request:  "TEZA",
			expected: "Hello TEZA",
		},
		{
			name:     "Hello('EKO')",
			request:  "EKO",
			expected: "Hello EKO",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

	fmt.Println("TestHelloTable done")
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Teza")
	}
}

func BenchmarkHelloSub(b *testing.B) {
	b.Run("Teza", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Teza")
		}
	})
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Eko")
		}
	})
}

func BenchmarkHelloTable(b *testing.B) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hello('TEZA')",
			request:  "TEZA",
			expected: "Hello TEZA",
		},
		{
			name:     "Hello('EKO')",
			request:  "EKO",
			expected: "Hello EKO",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(test.request)
			}
		})
	}
}
