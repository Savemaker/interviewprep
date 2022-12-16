package main

//Is Unique: Implement an algorithm to determine if a string has all unique characters.

func HasUniqueChars(input string) bool {
	hashset := make(map[rune]bool)
	for _, v := range input {
		_, ok := hashset[v]
		if ok {
			return false
		} else {
			hashset[v] = true
		}
	}
	return true
}
