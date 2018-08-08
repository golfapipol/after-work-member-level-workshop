package memberlevel

import "time"

const orderCountCondition = 8
const monthCondition = 6

type transaction struct {
	transactionId     int
	transactionPrice  float32
	transactionDate   time.Time
	transactionUserId string
}

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
func CheckUserID(userId int) string {

	if userId == 006 {
		return "Platinum"
	}
	return "Gold"
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
func GetLastSixMonthByUserId(userId string) []transaction {
	return []transaction{
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
		transaction{transactionPrice: 1000, transactionUserId: "006"},
	}
}
