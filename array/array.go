package array

// ContainsInt tests if an int is in a list of string.
func ContainsInt(l []int, s int) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}

	return false
}

// ContainsFloat tests if a float is in a list of string.
func ContainsFloat(l []float64, s float64) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}

	return false
}

// ContainsString tests if a string is in a list of string.
func ContainsString(l []string, s string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}

	return false
}
