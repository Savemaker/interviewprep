package main

/*
Check Permutation: Given two strings,
write a method to decide if one is a permutation of the other.
*/

func IsPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	hashmap := make(map[rune]bool)

	for _, v := range a {
		hashmap[v] = true
	}

	for _, v := range b {
		delete(hashmap, v)
	}

	return len(hashmap) == 0
}
