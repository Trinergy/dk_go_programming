package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, v := range s {
		n--
		runes[n] = v
	}
	return string(runes)
}

func isAnagram(s string) bool {
	reverse := reverse(s)
	return s == reverse
}

// commaBytes is the non recursive-way of adding commas to every third character
func commaBytes(s string) string {
	if len(s) <= 3 {
		return s
	}
	s = reverse(s)

	var buf bytes.Buffer
	counter := 0

	for i := range s {
		if counter == 3 {
			buf.WriteString(fmt.Sprintf(",%s", string(s[i])))
			counter = 1
		} else {
			buf.WriteString(fmt.Sprintf("%s", string(s[i])))
			counter++
		}
	}
	return reverse(buf.String())
}

func main() {
	w := "mmoomm"
	s := "1234567"
	fmt.Printf("reverse: %s\n", reverse(s))
	fmt.Printf("comma: %s\n", comma(s))
	fmt.Printf("commaBytes: %s\n", commaBytes(s))
	fmt.Printf("isAnagram: %t, word: %s\n", isAnagram(w), w)
}
