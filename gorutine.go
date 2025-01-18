package main

import "fmt"

// Функция обработки чисел
func fo(x int) int {
	return x * x
}

// Функция Merge2Channels
func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	for i := 0; i < n; i++ {
		// Чтение из обоих каналов
		x1 := <-in1
		x2 := <-in2

		// Обработка чисел функцией f и запись результата в out
		out <- f(x1) + f(x2)
	}
}

func main() {
	// Создаем каналы
	chan1 := make(chan int, 5) // Буфер для 5 элементов
	chan2 := make(chan int, 5)
	resultChan := make(chan int, 5)

	// Количество итераций
	n := 5

	// Заполняем каналы значениями
	for i := 1; i <= n; i++ {
		chan1 <- i
		chan2 <- i
	}

	// Запускаем функцию в отдельной горутине
	go Merge2Channels(fo, chan1, chan2, resultChan, n)

	// Читаем и выводим результаты
	for i := 0; i < n; i++ {
		fmt.Println(<-resultChan)
	}
}
