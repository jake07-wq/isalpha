package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		isUpper := char >= 'A' && char <= 'Z'
		isLower := char >= 'a' && char <= 'z'
		isDigit := char >= '0' && char <= '9'

		if !isUpper && !isLower && !isDigit {
			return false
		}
	}
	return true
}
