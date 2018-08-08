package memberlevel

const TMPID = "006"
const TMPLEVEL = "Platinum"

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

	if userId == TMPID {
		user.user_id = TMPID
		user.user_level = TMPLEVEL
		user.user_point = 0
	}

	return user
}