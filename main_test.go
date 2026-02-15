package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	arr := generateRandomElements(-1)
	assert.Equal(t, 0, len(arr))

	arr = generateRandomElements(0)
	assert.Equal(t, 0, len(arr))

	arr = generateRandomElements(1)
	assert.Equal(t, 1, len(arr))

	arr = generateRandomElements(100)
	assert.Equal(t, 100, len(arr))

	countUnique := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != arr[i+1] {
			countUnique++
		}
	}

	assert.Greater(t, countUnique, 50)
}

func TestMaximum(t *testing.T) {
	max := maximum(make([]int, 0))
	assert.Equal(t, 0, max)

	arr := []int{55}
	max = maximum(arr)
	assert.Equal(t, 55, max)

	arr = []int{4, 4, 4, 4, 4, 4, 4, 5}
	max = maximum(arr)
	assert.Equal(t, 5, max)

	arr = []int{5, 4, 4, 4, 4, 4, 4, 4}
	max = maximum(arr)
	assert.Equal(t, 5, max)

	arr = []int{4, 4, 4, 5, 4, 4, 4, 4}
	max = maximum(arr)
	assert.Equal(t, 5, max)

	arr = []int{-4, -4, -4, -5, -4, -4, 0, -4}
	max = maximum(arr)
	assert.Equal(t, 0, max)

	arr = []int{-4, -4, -4, -5, -4, -4, 0, 2}
	max = maximum(arr)
	assert.Equal(t, 2, max)

	arr = []int{4, 4, 4, 4, 4, 4, 4, 4}
	max = maximum(arr)
	assert.Equal(t, 4, max)

}

func TestMaxChunks(t *testing.T) {
	max := maximum(make([]int, 0))
	assert.Equal(t, 0, max)

	arr := []int{55}
	max = maximum(arr)
	assert.Equal(t, 55, max)

	arr = []int{4, 4, 4, 4, 4, 4, 4, 5}
	max = maximum(arr)
	assert.Equal(t, 5, max)

	arr = []int{5, 4, 4, 4, 4, 4, 4, 4}
	max = maximum(arr)
	assert.Equal(t, 5, max)

	arr = []int{4, 4, 4, 5, 4, 4, 4, 4}
	max = maximum(arr)
	assert.Equal(t, 5, max)

	arr = []int{-4, -4, -4, -5, -4, -4, 0, -4}
	max = maximum(arr)
	assert.Equal(t, 0, max)

	arr = []int{-4, -4, -4, -5, -4, -4, 0, 2}
	max = maximum(arr)
	assert.Equal(t, 2, max)

	arr = []int{4, 4, 4, 4, 4, 4, 4, 4}
	max = maximum(arr)
	assert.Equal(t, 4, max)
}
