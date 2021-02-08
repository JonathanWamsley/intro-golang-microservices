package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// initialization
	els := []int{9, 8, 7, 6, 5}

	// exection
	BubbleSort(els)

	// validation
	assert.NotNil(t, els)
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4]) 	
}

func TestBubbleSortBestCase(t *testing.T) {
	// initialization
	els := []int{5, 6, 7, 8, 9}

	// exection
	BubbleSort(els)

	// validation
	assert.NotNil(t, els)
	assert.EqualValues(t, []int{5, 6, 7, 8, 9}, els)
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4]) 	
}
func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
	

}

// Benchmark: test for efficiency

func getElement(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n-1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElement(t *testing.T) {
	els := getElement(5)
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 4, els[0])
	assert.EqualValues(t, 3, els[1])
	assert.EqualValues(t, 2, els[2])
	assert.EqualValues(t, 1, els[3])
	assert.EqualValues(t, 0, els[4])
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElement(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElement(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	els := getElement(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}


func BenchmarkSort10000(b *testing.B) {
	els := getElement(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort25000(b *testing.B) {
	els := getElement(25000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}


func BenchmarkSort25000(b *testing.B) {
	els := getElement(25000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}


func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElement(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}


func BenchmarkSort100000(b *testing.B) {
	els := getElement(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

// Benchmark results

// - bubble sort
// 1000	
	// 2246095	       546 ns/op
	// 30980	     38.199 ns/op
// 10000:
	// 196915	      5.747 ns/op
	//2458	    503.578 ns/op
// 25000
	// 41842	     25.839 ns/op
	// 828	   1.318.349 ns/op
// 100000:
	// 1	8.392.633.400 ns/op
	//201	   5.833.314 ns/op