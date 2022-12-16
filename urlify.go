/*
	URLify: Write a method to replace all spaces in a string with '%20'.
	You may assume that the string has sufficient space at the end to hold the additional characters,
	and that you are given the "true" length of the string.
	(Note:
	If implementing in Java,
	please use a character array so that you can perform this operation in place.)

	EXAMPLE
	Input: "Mr John Smith    ", 13 Output: "Mr%20John%20Smith"
*/

package main

func Urlify(input []rune, length int) []rune {
	c := 0
	for i, j := length-1, len(input)-1; i >= 0; i-- {
		if input[i] == ' ' {
			c++
			input[j] = '0'
			j--
			input[j] = '2'
			j--
			input[j] = '%'
			j--
		} else {
			input[j], input[i] = input[i], input[j]
			j--
		}
	}
	return input
}
