package fan_in_fan_out_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/d-jo/leetcode/fan_in_fan_out"
)

func TestMaxProcs(t *testing.T) {
	procs := runtime.GOMAXPROCS(0)

	fmt.Println(procs)
}

func TestPrimeFilter(t *testing.T) {
	input := int64(7)
	expected := true

	isPrime := true
	for i := int64(2); i < input; i++ {
		if input%i == 0 {
			isPrime = false
		}
	}

	if isPrime != expected {
		t.Errorf("Expected %t, but got %t", expected, isPrime)
	}
}

func TestPrimeFilter1(t *testing.T) {
	input := int64(6)
	expected := false

	isPrime := true
	for i := int64(2); i < input; i++ {
		if input%i == 0 {
			isPrime = false
		}
	}

	if isPrime != expected {
		t.Errorf("Expected %t, but got %t", expected, isPrime)
	}
}

func TestStartSerial(t *testing.T) {
	input := 100
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	actual := fan_in_fan_out.Start(input, 1)

	if len(actual) != len(expected) {
		t.Errorf("Expected %d, but got %d", len(expected), len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("(%d) Expected %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func TestStartParallel1(t *testing.T) {
	input := 100
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	actual := fan_in_fan_out.Start(input, 8)

	if len(actual) != len(expected) {
		t.Errorf("Expected %d, but got %d", len(expected), len(actual))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("(%d) Expected %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func TestStartParallel2(t *testing.T) {
	input := 100
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	actual := fan_in_fan_out.Start(input, 10)

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("(%d) Expected %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func BenchmarkPrimeFilter5Parallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(100, 8)
	}
}

func BenchmarkPrimeFilter5Parallel1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(1000, 8)
	}
}

func BenchmarkPrimeFilter5Parallel10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(10000, 8)
	}
}

func BenchmarkPrimeFilter10Parallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(100, 10)
	}
}

func BenchmarkPrimeFilter10Parallel1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(1000, 10)
	}
}

func BenchmarkPrimeFilter10Parallel10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(10000, 10)
	}
}

func BenchmarkPrimeFilter100Parallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(100, 100)
	}
}

func BenchmarkPrimeFilter100Parallel1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(1000, 100)
	}
}

func BenchmarkPrimeFilter100Parallel10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(10000, 100)
	}
}

func BenchmarkPrimeFilterSerial100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(100, 1)
	}
}

func BenchmarkPrimeFilterSerial1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(1000, 1)
	}
}

func BenchmarkPrimeFilterSerial10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fan_in_fan_out.StartNoSort(10000, 1)
	}
}
