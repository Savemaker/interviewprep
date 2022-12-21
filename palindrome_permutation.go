/*
	Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palin­drome.
	A palindrome is a word or phrase that is the same forwards and backwards.
	A permutation is a rearrangement of letters.
	The palindrome does not need to be limited to just dictionary words.
	EXAMPLE
	Input: Tact Coa
	Output: True (permutations: "taco cat", "atco cta", etc.)
*/

package main

import (
	"strings"
	"unicode"
)

func IsPermutationOfPalindrome(input string) bool {
	hashset := make(map[rune]bool)
	input = strings.ToLower(input)
	for _, r := range input {
		_, ok := hashset[r]
		if ok {
			delete(hashset, r)
		} else {
			if !unicode.IsSpace(r) {
				hashset[r] = true
			}
		}
	}
	lenght := len(hashset)
	if lenght == 0 || lenght == 1 {
		return true
	}
	return false
}
