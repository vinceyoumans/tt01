# Challenge TT01

package main

// All 3 functions return "random" words (in JSON format) from the given book
// Insert into output all words that are a palindrome (a word that reads the same backward as forward, e.g., madam)

func main() {
var output []string
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


