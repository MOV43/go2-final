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

// Создаём слайс со случайными положительными числами.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1_000_000) + 1
	}

	return data
}

// Находим максимум обычным перебором.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}

	return maxValue
}

// Находим максимум с помощью восьми горутин.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	results := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	// Начальное значение позволяет работать и с отрицательными числами.
	for i := 0; i < CHUNKS; i++ {
		results[i] = data[0]
	}

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Последняя часть забирает все оставшиеся элементы.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)

		go func(index int, part []int) {
			defer wg.Done()

			// У пустой части оставляем начальное значение.
			if len(part) > 0 {
				results[index] = maximum(part)
			}
		}(i, data[start:end])
	}

	// Ждём, пока все горутины запишут свои результаты.
	wg.Wait()

	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение последовательно")

	start := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение: %d\n", maxValue)
	fmt.Printf("Время поиска: %d мкс\n", elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)

	start = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение: %d\n", maxValue)
	fmt.Printf("Время поиска: %d мкс\n", elapsed)
}
