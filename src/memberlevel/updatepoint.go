package memberlevel

const tmpId = "006"
const tmpLevel = "Platinum"

type member struct {
	user_id string
	user_level string
	user_point int
}

func updatepoint(userId string, level string) int {
	var user member

	user = getuserdata(userId)

	//freepoint = GetFreePoint(userId, Level)
	freepoint := 800

	user.user_point = user.user_point + freepoint

	return user.user_point
}

func getuserdata(userId string) member {
	var user member

	if userId == tmpId {
		user.user_id = tmpId
		user.user_level = tmpLevel
		user.user_point = 0
	}

	return user
}