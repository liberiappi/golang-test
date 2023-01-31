package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRafi(t *testing.T){
	result := HelloWorld("Rafi")

	assert.Equal(t, "Hello Rafi", result, "Result must be 'Hello Rafi'")


	fmt.Println("TestHelloWorldRafi done")
}

func TestHelloWorldAlfatih(t *testing.T){
	result := HelloWorld("Alfatih")

	if result != "Hello Alfatih" {
		t.Fatal("Result must be 'Hello Alfatih'")
	}

	fmt.Println("TestHelloWorldAlfatih done")
}


func TestHelloAssert(t *testing.T){
	result := HelloWorld("Rapi")

	assert.Equal(t, "Hello Rapi", result, "Result must be 'Hello Rapi'")
}

func TestHelloRequire(t *testing.T){
	result := HelloWorld("Rapi")

	require.Equal(t, "Hello Rapi", result, "Result must be 'Hello Rapi'")
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Cannot run on windows")
	}

	result := HelloWorld("Rapi")

	require.Equal(t, "Hello Rapi", result, "Result must be 'Hello Rapi'")
}

func TestMain(m *testing.M){
	fmt.Println("Before")

	m.Run()

	fmt.Println("After")
}

func TestSubTest(t *testing.T) {
	t.Run("Rafi", func(t *testing.T){
		result := HelloWorld("Rafi")
		require.Equal(t, "Hello Rafi", result)
	})
	t.Run("Alfatih", func(t *testing.T){
		result := HelloWorld("Alfatih")
		assert.Equal(t, "Hello Alfatih", result)
	})
}

func TestTableTest(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Rafi",
			request: "Rafi",
			expected: "Hello Rafi",
		},
		{
			name: "Alfatih",
			request: "Alfatih",
			expected: "Hello Alfatih",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++{
		HelloWorld("Rafi")
	}
}

func BenchmarkHelloWorldAlfatih(b *testing.B) {
	for i := 0; i < b.N; i++{
		HelloWorld("Alfatih")
	}
}

func BenchmarkSubHello(b *testing.B) {
	b.Run("Rafi", func(b *testing.B){
		for i := 0; i < b.N; i++{
			HelloWorld("Rafi")
		}
	})
	b.Run("Alfatih", func(b *testing.B){
		for i := 0; i < b.N; i++{
			HelloWorld("Alfatih")
		}
	})
}

func BenchmarkTableHello(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	} {
		{
			name: "Rafi",
			request: "Rafi",
		},
		{
			name: "Alfatih",
			request: "Alfatih",
		},
		{
			name: "Muhammad",
			request: "Muhammad",
		},
		{
			name: "Hani",
			request: "Hani",
		},
		{
			name: "Ridha",
			request: "Ridha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i ++{
				HelloWorld(benchmark.request)
			}
		})
	}
}
