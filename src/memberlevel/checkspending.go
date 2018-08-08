package memberlevel

func CheckSpending(spending int) bool {
	if spending > 1000 {
		return true
	}
	return false
}
