package memberlevel

type Transaction struct {
	product string
	price float64
	date string
}
type Transactions []Transaction
const price = 1000
func filterTranscationBySpending(transactions Transactions) (Transactions) {

	filterTransactions := Transactions{}
	for i:=0; i< len(transactions);i++ {
		if transactions[i].price > price {
			//filterTransactions[i] transaction:=	transactions[i]
			filterTransactions = append(filterTransactions, transactions[i])
		}
	}
	
	return filterTransactions
}
