package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// All 3 functions return "random" words (in JSON format) from the given book
// Insert into output all words that are a palindrome (a word that reads the same backward as forward, e.g., madam)

type book struct {
	Page  int
	Word  string
	isPal bool
}

var TBOOKS [][]book

func main() {

	TBOOKS = append(TBOOKS, Step002(BookA()))

	TBOOKS = append(TBOOKS, Step002(BookB()))

	TBOOKS = append(TBOOKS, Step002(BookC()))

	fmt.Println(TBOOKS)

}

func Step002(str []byte) []book {
	// converts []bytes to TBOOK
	var TBOOK []book
	err := json.Unmarshal(str, &TBOOK)
	if err != nil {
		fmt.Println("error:", err)
	}

	for i, B := range TBOOK {
		fmt.Println(i, B.Word)

		fmt.Println(isPalendrome(B.Word))
		if isPalendrome(B.Word) {
			TBOOK[i].isPal = true
		} else {
			TBOOK[i].isPal = false
		}
	}

	return TBOOK
}

func BookA() []byte {
	return []byte(`[{"page": 2, "word": "Racecar"}, {"page": 13, "word": "Speed"}, {"page": 24, "word": "rotoR"} ]`)
}

func BookB() []byte {
	return []byte(`[{"page": 4, "word": "Mario"}, {"page": 11, "word": "Level"}, {"page": 39, "word": "Luigi"} ]`)
}

func BookC() []byte {
	return []byte(`[{"page": 8, "word": "Madam"}, {"page": 19, "word": "Tornado"}, {"page": 43, "word": "Radar"} ]`)
}

func isPalendrome(strr string) bool {

	str := strings.ToUpper(strr)

	result := []byte{}

	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}

	return str == string(result)

}
