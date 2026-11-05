package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	sizes := []int{-5, 0, 1, 100}

	for _, size := range sizes {
		data := generateRandomElements(size)

		expectedSize := size
		if size < 0 {
			expectedSize = 0
		}

		if len(data) != expectedSize {
			t.Errorf(
				"Размер %d: получили длину %d, ожидали %d",
				size, len(data), expectedSize,
			)
		}

		for _, value := range data {
			if value <= 0 {
				t.Errorf("Получили неположительное число: %d", value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{5}, 5},
		{"максимум в начале", []int{9, 2, 3}, 9},
		{"максимум в середине", []int{2, 9, 3}, 9},
		{"максимум в конце", []int{2, 3, 9}, 9},
		{"одинаковые числа", []int{7, 7, 7}, 7},
		{"нули", []int{0, 0, 0}, 0},
		{"отрицательные числа", []int{-5, -2, -8}, -2},
		{"восемь элементов", []int{1, 8, 3, 4, 5, 6, 7, 2}, 8},
		{"максимум в остатке", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maximum(test.data)

			if result != test.want {
				t.Errorf(
					"maximum: получили %d, ожидали %d",
					result, test.want,
				)
			}

			result = maxChunks(test.data)

			if result != test.want {
				t.Errorf(
					"maxChunks: получили %d, ожидали %d",
					result, test.want,
				)
			}
		})
	}
}
