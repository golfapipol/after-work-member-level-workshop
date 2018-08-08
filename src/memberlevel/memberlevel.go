package memberlevel

import "testing"

const orderCountCondition = 8
const monthCondition = 6

func CheckCountOrder(order int, month int) bool {
	if CheckOrderPerMonth(month) {
		if order >= orderCountCondition {
			return true
		}
	}
	return false
}

func CheckOrderPerMonth(month int) bool {
	if month <= monthCondition {
		return true
	}
	return false
}

func CheckSpending(spending int) bool {
	if spending > 1000 {
		return true
	}
	return false
}

func CheckUserInRankByID(userId string) bool {
	id := "006"
	if userId == id {
		return true
	}
	return false
}

type Transaction struct {
	product string
	price   float64
	date    string
}
type Transactions []Transaction

const price = 1000

func filterTranscationBySpending(transactions Transactions) Transactions {

	filterTransactions := Transactions{}
	for i := 0; i < len(transactions); i++ {
		if transactions[i].price > price {
			filterTransactions = append(filterTransactions, transactions[i])
		}
	}

	return filterTransactions
}
func Test_GetLastSixMonthByUserId_Input_006_Should_Be_Transaction_Within_Six_Month(t *testing.T) {
	expectedCountTransactionNumber := 8
	userId := "006"

	actualTransactions := GetLastSixMonthByUserId(userId)

	if len(actualTransactions) != expectedCountTransactionNumber {
		t.Errorf("Error, the transaction should be (%v) but it is (%v)", expectedCountTransactionNumber, len(actualTransactions))
	}
}
