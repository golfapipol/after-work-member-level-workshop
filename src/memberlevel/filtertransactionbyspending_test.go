package memberlevel

import "testing"

func Test_filterTranscationBySpending_input_fan_1200_10feb2018_Should_Be_fan_1200_10feb2018(t *testing.T) {

	transactions := Transactions{
		Transaction{product: "พัดลม", price: 1200.00, date: "10 Feb 2018"},
		Transaction{product: "กางเกง", price: 800.00, date: "2 Mar 2018"},
		Transaction{product: "มือถือ", price: 1000.00, date: "7 Jul 2018"},
	}
	expectedTransactions := Transactions{
		Transaction{product: "พัดลม", price: 1200.00, date: "10 Feb 2018"},
	}
	expectedTransactionCount := len(expectedTransactions)

	actualTransactions := filterTranscationBySpending(transactions)

	if expectedTransactionCount != len(actualTransactions) {
		t.Errorf("Expected Count %v but got it %v", expectedTransactionCount, len(actualTransactions))

	}
	if expectedTransactions[0] != actualTransactions[0] {
		t.Errorf("ExpectedProduct %s but got it %s", expectedTransactions[0].product, actualTransactions[0].product)
	}
	// if expectedProduct != actualProduct {
	// 	t.Errorf("ExpectedProduct %s but got it %s", expectedProduct, actualProduct)
	// }
	// if expectedPrice != actualPrice {
	// 	t.Errorf("ExpectedProduct %v but got it %v", expectedPrice, actualPrice)
	// }
	// if expectedDate != actualDate {
	// 	t.Errorf("ExpectedProduct %v but got it %v", expectedDate, actualDate)
	// }
}
