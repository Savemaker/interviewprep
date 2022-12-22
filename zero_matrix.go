package main

func ZeroMatrix(input [][]int) [][]int {
	type point struct {
		i, j int
	}

	hashmap := make(map[int]point, 0)

	key := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				hashmap[key] = point{i: i, j: j}
				key++
			}
		}
	}

	for _, v := range hashmap {
		for i := v.i - 1; i >= 0; i-- {
			input[i][v.j] = 0
		}
		for i := v.i + 1; i < len(input); i++ {
			input[i][v.j] = 0
		}
		for j := v.j - 1; j >= 0; j-- {
			input[v.i][j] = 0
		}
		for j := v.j + 1; j < len(input[v.i]); j++ {
			input[v.i][j] = 0
		}
	}
	return input
}
