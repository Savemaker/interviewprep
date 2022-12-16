package main

/*
Check Permutation: Given two strings,
write a method to decide if one is a permutation of the other.
*/

func IsPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	hashmap := make(map[rune]int)

	for _, v := range a {
		hashmap[v] += 1
	}

	for _, v := range b {
		count, ok := hashmap[v]
		if ok {
			if count-1 == 0 {
				delete(hashmap, v)
			} else {
				hashmap[v] -= 1
			}
		}
	}
	return len(hashmap) == 0
}
