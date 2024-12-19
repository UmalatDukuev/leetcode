package main

import (
	"fmt"
)

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
	fmt.Println(spaces)

	result_strings := make([]string, len(spaces))
	substr := ""

	for i := 0; i < len(spaces)-1; i++ {

		j := int(spaces[i])
		for j < int(spaces[i+1])-1 {
			// fmt.Print(words[j])
			// fmt.Print(" ")

			substr += words[j] + " "
			j++
		}
		// fmt.Println("")

		substr += words[j]
		fmt.Println(substr)
		substr = "\"" + substr + "\""
		result_strings[i] = substr
		substr = ""
	}

	fmt.Println(result_strings)

	return []string{}
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20
	fmt.Println(fullJustify(words, maxWidth))
}
