package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Yohan",
			request: "Yohan",
		},
		{
			name:    "Nabila",
			request: "Nabila",
		},
		{
			name:    "Fathar",
			request: "Fathar",
		},
		{
			name:    "Taqy",
			request: "Taqy",
		},
		{
			name:    "Yohan Apriyandi",
			request: "Yohan Apriyandi",
		},
		{
			name:    "Fithriya Nabilah",
			request: "Fithriya Nabilah",
		},
		{
			name:    "Fathar Dhabit Adz-Dzaky",
			request: "Fathar Dhabit Adz-Dzaky",
		},
		{
			name:    "Fariq Taqy Abqary",
			request: "Fariq Taqy Abqary",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWord(benchmark.request)
			}
		})
	}
}

func BenchmarkSubHelloWord(b *testing.B) {
	b.Run("Yohan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWord("Yohan")
		}
	})

	b.Run("Nabilah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWord("Nabilah")
		}
	})
}

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("aa")
	}
}

func BenchmarkHelloWordFathar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("ade")
	}
}

func TestTableHelloWord(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Yohan",
			request:  "Yohan",
			expected: "Hello Yohan",
		},
		{
			name:     "Nabila",
			request:  "Nabila",
			expected: "Hello Nabila",
		},
		{
			name:     "Fathar",
			request:  "Fathar",
			expected: "Hello Fathar",
		},
		{
			name:     "Taqy",
			request:  "Taqy",
			expected: "Hello Taqy",
		},
		{
			name:     "Fithriya",
			request:  "Fithriya",
			expected: "Hello Fithriya",
		},
		{
			name:     "Apriyandi",
			request:  "Apriyandi",
			expected: "Hello Apriyandi",
		},
		{
			name:     "Fariq",
			request:  "Fariq",
			expected: "Hello Fariq",
		},
		{
			name:     "Dhabit",
			request:  "Dhabit",
			expected: "Hello Dhabit",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWord(test.request)
			require.Equal(t, test.expected, result, "Berhasil di eksekusi")
		})
	}
	fmt.Println("Berhasil diksekusi testnya")
}

func TestSubTest(t *testing.T) {
	t.Run("Yohan", func(t *testing.T) {
		result := HelloWord("Yohan")
		require.Equal(t, "Hello Yohan", result, "Result must be 'Hello Yohan'")
	})
	t.Run("Apriyandi", func(t *testing.T) {
		result := HelloWord("Apriyandi")
		require.Equal(t, "Hello Apriyandi", result, "Result must be 'Hello Apriyandi'")
	})
}

func TestMain(m *testing.M) {
	// after
	fmt.Println("Before unit test")

	m.Run()

	// before
	fmt.Println("After unit test")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on mac OS")
	}
	result := HelloWord("Yohan")
	assert.Equal(t, "Hello Yohan", result, "Result must be 'Hello Yohan'")
	fmt.Println("Test with t.Skip done")
}

func TestHelloWordAssert(t *testing.T) {
	//go get github.com/stretchr/testify/assert@v1.10.0 intall it if u use assert
	result := HelloWord("Yohan")
	assert.Equal(t, "Hello Yohan", result, "Result must be 'Hello Yohan'")
	fmt.Println("Test with assertion done")
}

func TestHelloWordRequire(t *testing.T) {
	//go get github.com/stretchr/testify/assert@v1.10.0 intall it if u use assert
	result := HelloWord("Yohan")
	require.Equal(t, "Hello Yohan", result, "Result must be 'Hello Yohan'")
	fmt.Println("Test with require done")
}

func TestHelloWord(t *testing.T) {
	result := HelloWord("Yohan")
	if result != "Hello Yohan" {
		//panic("result it not Hello Yohan")
		//t.Fail()
		//t.Error() bisa diberikan argumen sehingga ketika ada error lebih mudah dibaca
		//t.Error akan memanggil t.fail()
		t.Error("result must be Hello Yohan")
	}
	fmt.Println("Test hello word done")
}

func TestHelloFathar(t *testing.T) {
	result := HelloWord("Fathar")
	if result != "Hello Fathar" {
		//t.FailNow()
		//t.Fatal() bisa diberikan argumen sehingga ketika ada error lebih mudah dibaca
		//t.Error akan memanggil t.FailNow()
		t.Fatal("result must be Hello fathar")
	}
	fmt.Println("Test hello word fathar done")
}
