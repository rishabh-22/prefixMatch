package matcher

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func getMatchingIndex(prefixes []string, string string) int {
	lo := 0
	hi := len(prefixes) - 1
	if len(prefixes) == 0 {
		return -1 // return if the prefix array is empty
	}
	for lo <= hi {
		mid := (lo + hi) / 2
		size := Min(len(prefixes[mid]), len(string))

		if prefixes[mid] < string[:size] {
			lo = mid + 1
		} else if prefixes[mid] > string[:size] {
			hi = mid - 1
		} else {
			if mid+1 >= len(prefixes) {
				return mid
			}
			size = Min(len(prefixes[mid+1]), len(string))
			if prefixes[mid+1] != string[:size] {
				return mid
			}
			lo = mid + 1
		}
	}
	return -1
}

func LongestPrefix(prefixes []string, string string) string {
	idx := getMatchingIndex(prefixes, string)
	if idx == -1 {
		return "No prefix found"
	}
	return prefixes[idx]
}
