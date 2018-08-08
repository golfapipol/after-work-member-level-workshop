package memberlevel

func CheckUserID(userId int) string {

	if userId == 006 {
		return "Platinum"
	}
	return "Gold"

}
