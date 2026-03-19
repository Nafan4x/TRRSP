package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func clearInputStream(reader *bufio.Reader) {
	// Читаем всё до новой строки, чтобы очистить буфер
	reader.ReadString('\n')
}

func getValidatedInt(reader *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число.")
			continue
		}

		return value
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	N := getValidatedInt(reader, "Введите порядок матрицы N: ")

	if N <= 0 {
		fmt.Println("Ошибка: порядок матрицы должен быть положительным числом.")
		return
	}

	// Создаем матрицу N x N
	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	fmt.Println("Введите элементы матрицы:")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for {
				fmt.Printf("Элемент [%d][%d]: ", i+1, j+1)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					continue
				}

				input = strings.TrimSpace(input)
				value, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Ошибка: введите целое число.")
					continue
				}

				matrix[i][j] = value
				break
			}
		}
	}

	fmt.Println("\nИсходная матрица:")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}

	// Поиск минимального элемента на главной и побочной диагоналях
	minElement := int(^uint(0) >> 1) // Максимальное значение int
	for i := 0; i < N; i++ {
		if matrix[i][i] < minElement {
			minElement = matrix[i][i]
		}
		if matrix[i][N-1-i] < minElement {
			minElement = matrix[i][N-1-i]
		}
	}

	fmt.Printf("\nМинимальный элемент на диагоналях: %d\n", minElement)

	// Заполнение элементов выше главной диагонали минимальным элементом
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			matrix[i][j] = minElement
		}
	}

	// Вывод измененной матрицы (добавил для наглядности)
	fmt.Println("\nИзмененная матрица (элементы выше главной диагонали заменены):")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}
