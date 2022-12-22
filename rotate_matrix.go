/*
	Rotate Matrix:
	Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
	write a method to rotate the image by 90 degrees. Can you do this in place?

	omg for the first time in like 4 years and God knows how many tries I did it without reaching for help.
	Hell yeah
*/

package main

func RotateMatrix(data *[][]int) [][]int {
	input := *data
	size := len(input)

	for i := 0; i < size/2; i++ {
		for j := i; j < size-1-i; j++ {
			t := input[i][j]
			input[i][j] = input[size-1-j][i]
			input[size-1-j][i] = input[size-1-i][size-1-j]
			input[size-1-i][size-1-j] = input[j][size-1-i]
			input[j][size-1-i] = t
		}
	}
	return input
}
