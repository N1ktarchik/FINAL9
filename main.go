package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	arr := make([]int, size)

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	for i := 0; i < size; i++ {
		arr[i] = rnd.Int()
	}

	return arr
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	maxElem := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > maxElem {
			maxElem = data[i]
		}
	}

	return maxElem
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < 8 {
		return maximum(data)
	}

	lenArr := len(data) / 8
	resultArr := make([]int, 8)

	wg := sync.WaitGroup{}

	for i := 0; i < 8; i++ {

		startIndex := i * lenArr
		var lastIndex int

		if i == 7 {
			lastIndex = len(data)
		} else {
			lastIndex = startIndex + lenArr
		}

		wg.Add(1)
		go func(mas []int, index int) {
			defer wg.Done()

			maxElem := mas[0]

			for _, v := range mas[1:] {
				if v > maxElem {
					maxElem = v
				}
			}

			resultArr[index] = maxElem

		}(data[startIndex:lastIndex], i)

	}

	wg.Wait()

	maxElem := resultArr[0]

	for i := 1; i < 8; i++ {
		if resultArr[i] > maxElem {
			maxElem = resultArr[i]
		}
	}

	return maxElem

}

func main() {
	fmt.Printf("Генерируем %d целых чисел ", SIZE)
	arr := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	initTime := time.Now()
	max := maximum(arr)
	elapsed := time.Since(initTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков ", CHUNKS)

	initTime = time.Now()
	max = maxChunks(arr)
	elapsed = time.Since(initTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
