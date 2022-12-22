/*
	String Compression:
	Implement a method to perform basic string compression using the counts of repeated characters.
	For example, the string aabcccccaaa would become a2b1c5a3.
	If the "compressed" string would not become smaller than the original string,
	your method should return the original string.
	You can assume the string has only uppercase and lowercase letters (a - z).
*/

package main

import (
	"strconv"
	"strings"
)

func CompressString(input string) string {
	type compressed struct {
		letter byte
		times  int
	}

	resSlice := make([]*compressed, 0)
	position := 0
	inputLen := len(input)
	var currentChar byte

	for i := 0; i < inputLen; i++ {
		if len(resSlice) == 0 {
			resSlice = append(resSlice, &compressed{letter: input[i], times: 1})
			currentChar = input[i]
		} else {
			if input[i] == currentChar {
				query := resSlice[position]
				query.times += 1
			} else {
				position++
				currentChar = input[i]
				resSlice = append(resSlice, &compressed{letter: input[i], times: 1})
			}
		}
	}

	var resBuilder strings.Builder

	for i := range resSlice {
		qeury := resSlice[i]
		resBuilder.WriteByte(qeury.letter)
		resBuilder.WriteString(strconv.Itoa(qeury.times))
	}

	result := resBuilder.String()
	if len(result) > inputLen {
		return input
	}
	return result
}
