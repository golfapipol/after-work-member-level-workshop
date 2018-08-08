package memberlevel

import "time"

type transaction struct {
	transactionId     int
	transactionPrice  float32
	transactionDate   time.Time
	transactionUserId string
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
