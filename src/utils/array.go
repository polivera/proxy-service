package utils

func InStringArray(needle string, highStack []string) bool {
	for _, item := range highStack {
		if item == needle {
			return true
		}
	}
	return false
}
