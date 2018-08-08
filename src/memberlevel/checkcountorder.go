package memberlevel

const orderCountCondition = 8
const monthCondition = 6

func CheckCountOrder(order int,month int) bool{
	if CheckOrderPerMonth(month){
		if order >= orderCountCondition {
			return true
		}
	}
	return false
}

func CheckOrderPerMonth(month int) bool{
	if month <= monthCondition {
		return true
	}
	return false
}