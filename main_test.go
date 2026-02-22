package main

import (
	"testing"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {

	length := []int{-1, 0, 8, 10, 100}
	expected := []int{0, 0, 8, 10, 100}

	for i := 0; i < len(length); i++ {
		mas := generateRandomElements(length[i])

		if len(mas) != expected[i] {
			t.Errorf("generate random element error. Length: %d , Expected: %d , Result: %d", length[i], expected[i], len(mas))
		}
	}

}

func TestMaximum(t *testing.T) {
	tests := [][]int{
		{},
		{13},
		{3, 2, 4, 7, 3, 4, 1},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 1, 1, 1, 2, 1, 1, 1, 1},
	}

	expected := []int{0, 13, 7, 8, 9, 10, 2}

	for i := 0; i < len(expected); i++ {
		maxElem := maximum(tests[i])

		if maxElem != expected[i] {
			t.Errorf("func maximum error. Mas: %v , Result: %d , Expected %d", tests[i], maxElem, expected[i])
		}
	}

}

func TestMaxChunks(t *testing.T) {
	tests := [][]int{
		{},
		{13},
		{3, 2, 4, 7, 3, 4, 1},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 1, 1, 1, 2, 1, 1, 1, 1},
	}

	expected := []int{0, 13, 7, 8, 9, 10, 2}

	for i := 0; i < len(expected); i++ {
		maxElem := maxChunks(tests[i])

		if maxElem != expected[i] {
			t.Errorf("func maxChunks error. Mas: %v , Result: %d , Expected %d", tests[i], maxElem, expected[i])
		}
	}
}
