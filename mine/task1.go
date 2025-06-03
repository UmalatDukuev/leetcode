package main

func main() {

}

// мерж отсортированных слайсов

// func mergeSlices(sl1, sl2 []int) []int {
// 	len1 := len(sl1)
// 	len2 := len(sl2)
// 	resLen := len1 + len2
// 	result := make([]int, 0, resLen)
// 	i, j := 0, 0
// 	for i < len1 && j < len2 {
// 		if sl1[i] <= sl2[j] {
// 			result = append(result, sl1[i])
// 			i++
// 		} else {
// 			result = append(result, sl2[j])
// 			j++

// 		}
// 	}
// 	result = append(result, sl2[j:]...)
// 	result = append(result, sl1[i:]...)
// 	return result
// }

// добавление 10-ью горутинами значений в слайс безопасно

// func main() {
// 	var (
// 		wg    sync.WaitGroup
// 		mutex sync.Mutex
// 		slice []int
// 	)
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)

// 		go func(i_ int) {
// 			mutex.Lock()
// 			defer mutex.Unlock()
// 			slice = append(slice, i_)
// 			wg.Done()
// 		}(i)

// 	}
// 	wg.Wait()
// 	fmt.Println(slice)
// }
