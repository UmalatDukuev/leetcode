package main

import (
	"fmt"
	"sort"
	"strings"
)

type Key struct {
	Even string
	Odd  string
}

// func findSimilar(strs []string) int {
// 	numOfStrs := len(strs)
// 	countMap := make(map[Key]int)
// 	for i := 0; i < numOfStrs; i++ {
// 		strLen := len(strs[i])
// 		// key, exists := countMap[strs[i]]
// 		even, odd := splitEvenOdd(s)
// 		if !exists {
// 			for j := 0; j < strLen; j++ {
// 				if j%2 == 0 {
// 					key.Even += string(strs[i][j])
// 				} else {
// 					key.Odd += string(strs[i][j])
// 				}
// 			}
// 			countMap[strs[i]] = key
// 		}
// 	}
// 	fmt.Println(countMap)
// 	return 4
// }

func findSimilar(strs []string) int {
	// Мапа для подсчета строк с одинаковыми ключами
	countMap := make(map[Key]int)

	for _, s := range strs {
		// Разделяем символы строки на четные и нечетные
		even, odd := splitEvenOdd(s)

		// Сортируем символы для создания уникального ключа
		sort.Strings(even)
		sort.Strings(odd)

		// Формируем ключ
		key := Key{
			Even: strings.Join(even, ""),
			Odd:  strings.Join(odd, ""),
		}

		// Увеличиваем счетчик для данного ключа
		countMap[key]++
	}
	fmt.Println(countMap)
	// Подсчитываем количество пар
	totalPairs := 0
	for _, count := range countMap {
		// Число пар в группе: count * (count - 1) / 2
		totalPairs += count * (count - 1) / 2
	}

	return totalPairs
}

func splitEvenOdd(s string) ([]string, []string) {
	var even, odd []string
	for i, char := range s {
		if i%2 == 0 {
			even = append(even, string(char))
		} else {
			odd = append(odd, string(char))
		}
	}
	return even, odd
}

func main() {
	// in := bufio.NewReader(os.Stdin)
	// out := bufio.NewWriter(os.Stdout)
	// defer out.Flush()

	// var t int
	// fmt.Fscan(in, &t)
	// result := make([]int, 0)
	// for i := 0; i < t; i++ {
	// 	var n int
	// 	fmt.Fscan(in, &n)
	// 	strs := make([]string, n)
	// 	for j := 0; j < n; j++ {
	// 		var s string
	// 		fmt.Fscan(in, &s)
	// 	}
	// 	res := findSimilar(strs)
	// 	result = append(result, res)
	// }

	result := make([]int, 0)
	strs := []string{"ababa", "ababa", "ababa", "avaca"}
	res := findSimilar(strs)
	result = append(result, res)
	fmt.Println(result)
}
