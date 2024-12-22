package main

import (
	"fmt"
	"strings"
)

func spacing(str []string, maxWidth int, wordsLen int) string {
	maxSpace := maxWidth - wordsLen
	if len(str) == 1 {
		return str[0] + strings.Repeat(" ", maxSpace)
	}
	if len(str) == 2 {
		return str[0] + strings.Repeat(" ", maxSpace) + str[1]
	}
	gaps := len(str) - 1
	baseSpace := maxSpace / gaps
	extraSpaces := maxSpace % gaps

	result := ""

	for i := 0; i < len(str); i++ {
		if i < extraSpaces {
			result += str[i] + strings.Repeat(" ", baseSpace) + " "
		} else {
			result += str[i] + strings.Repeat(" ", baseSpace)
		}

	}
	return result
}

func leftJustify(str []string, maxWidth int, wordsLen int) string {
	maxSpace := maxWidth - wordsLen
	if len(str) == 1 {
		return str[0] + strings.Repeat(" ", maxSpace)
	}
	gaps := len(str) - 1
	result := ""
	for i := 0; i <= gaps; i++ {
		result += str[i] + strings.Repeat(" ", 1)
		maxSpace--
	}
	result += strings.Repeat(" ", maxSpace)
	return result
}

func fullJustify(words []string, maxWidth int) []string {
	spaces := make([]int, 1)
	wordsLen := len(words)
	strLen := 0

	for i := 0; i < wordsLen; i++ {
		strLen += len(words[i]) + 1
		if strLen > maxWidth+1 {
			strLen = 0
			spaces = append(spaces, i)
			i--
		}
	}
	spaces = append(spaces, wordsLen) // массив с индексами слов, после которых нужен перенос строки
	result := make([]string, 0)
	for i := 0; i < len(spaces)-1; i++ {
		subStrs := make([]string, 0)
		wordsLen := 0
		j := int(spaces[i])
		for j < int(spaces[i+1]) {
			wordsLen += len(words[j])
			subStrs = append(subStrs, words[j])
			j++
		} // в этом цикле группируем в массив строк слова одной строки
		if i != len(spaces)-2 {
			substr := spacing(subStrs, maxWidth, wordsLen)
			result = append(result, substr)
			fmt.Println(substr)
		} else {
			substr := leftJustify(subStrs, maxWidth, wordsLen)
			result = append(result, substr)
			fmt.Println(substr)
		} // последняя строка обрабатывается иначе
	}

	return []string{}
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	fullJustify(words, maxWidth)
}
