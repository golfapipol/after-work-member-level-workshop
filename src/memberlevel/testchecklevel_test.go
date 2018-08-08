package memberlevel_test

func CheckLevel(level string) int {
	freePoint := make(map[string]int)
	freePoint["Daimond"] = 1500
	freePoint["Platinum"] = 300
	freePoint["Silver"] = 200
	freePoint["Gold"] = 100

	return freePoint[level]
}
