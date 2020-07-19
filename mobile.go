package zhong

// CheckMobile checks whether str is a valid mobile
func CheckMobile(str string) bool {
	const digits = 11
	if len(str) != digits {
		return false
	}
	if str[0] != '1' {
		return false
	}
	if str[1] < '3' || str[1] > '9' {
		return false
	}
	for i := 2; i < digits; i++ {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	return true
}
