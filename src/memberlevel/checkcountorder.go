package memberlevel

const orderIdCondition = 8

func CheckCountOrder(orderId int) bool{

	if orderId >= orderIdCondition {
		return true
	}
	return false
}