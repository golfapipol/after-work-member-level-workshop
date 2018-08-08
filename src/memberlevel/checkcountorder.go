package memberlevel

const orderCondition = 8

func CheckCountOrder(order int) bool {

	if order >= orderCondition {
		return true
	}
	return false
}
