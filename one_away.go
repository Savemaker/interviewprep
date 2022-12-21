package main

func OneAway(a, b string) bool {
	isSubPossible := false

	lenA := len(a)
	lenB := len(b)

	dif := lenA - lenB
	if dif < 0 {
		dif *= -1
	}

	if dif > 1 {
		return false
	} else if dif == 0 {
		isSubPossible = true
	}

	var str, cmp string
	var strLen, cmpLen int

	if lenA >= lenB {
		str = a
		strLen = lenA
		cmp = b
		cmpLen = lenB
	} else {
		str = b
		strLen = lenB
		cmp = a
		cmpLen = lenA
	}

	changes := 0
	for i, j := 0, 0; i < strLen; i++ {
		if isSubPossible {
			if str[i] != cmp[j] {
				changes += 1
				if changes > 1 {
					return false
				}
			}
			j++
		} else {
			if j >= cmpLen || str[i] != cmp[j] {
				changes += 1
				if changes > 1 {
					return false
				}
			} else {
				j++
			}
		}
	}
	return changes == 0 || changes == 1
}
