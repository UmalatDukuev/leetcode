package main

import (
	"fmt"
	"strconv"
)

func checkPowTwo(num int) (bool, int) {
	strin := ""
	for num > 0 {
		strin += strconv.Itoa(num % 2)
		num /= 2
	}
	strLen := len(strin) - 1
	for i := 0; i < strLen; i++ {
		if i != strLen && strin[i] == '1' {
			return false, -1
		}
	}
	return true, strLen
}

func findThree(num int) int {
	var mas []int
	cnt := 0
	for i := num; i > 0; i-- {
		flag, _ := checkPowTwo(i)
		if flag == true && len(mas) < 3 && cnt+i <= num {
			cnt += i
			mas = append(mas, i)
		}
	}
	if len(mas) != 3 {
		return -1
	}
	return cnt
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		fmt.Println(findThree(a))
	}
}

/*
В самый что ни на есть обычный день Борис решил, что на протяжении следующих  дней он будет покупать своей маме букет ровно из трех цветов. Там, где живет Борис, существует всего лишь один магазин цветов, но зато в нем широкий ассортимент: для каждого  от  до  в магазине есть ровно один уникальный цветок, стоящий  бурлей. К тому же в магазин ежедневно довозят цветы, которые были выкуплены.
В -й день у Бориса есть  бурлей, которые он готов потратить на букет. Борис хочет купить как можно более дорогой букет. Для каждого из  дней определите, за какую стоимость Борис купит букет, или сообщите, что на его деньги невозможно купить никакой букет из трех цветов.
*/
