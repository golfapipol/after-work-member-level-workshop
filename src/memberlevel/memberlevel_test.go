package memberlevel

import "testing"

func Test_CheckCountOrder_Input_8_should_Be_True(t *testing.T) {
	expected := true
	order := 8
	month := 6

	actualResult := CheckCountOrder(order, month)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)
	}

}

func Test_CheckCountOrder_Input_8_should_Be_False(t *testing.T) {
	expected := false
	order := 7
	month := 7

	actualResult := CheckCountOrder(order, month)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)
	}

}

func Test_CheckOrderPerMonth_Input_6_should_Be_True(t *testing.T) {
	expected := true
	month := 6

	actualResult := CheckOrderPerMonth(month)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, month)
	}

}

func Test_CheckOrderPerMonth_Input_7_should_Be_False(t *testing.T) {
	expected := false
	month := 7

	actualResult := CheckOrderPerMonth(month)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, month)
	}

}

func Test_CheckSpending_Input_1200_Should_Be_True(t *testing.T) {
	spending := 1200
	expected := true

	actual := CheckSpending(spending)

	if expected != actual {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}

func Test_CheckSpending_Input_800_Should_Be_False(t *testing.T) {
	spending := 800
	expected := false

	actual := CheckSpending(spending)

	if expected != actual {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}

func Test_CheckUserInRankByID_Input_006_Should_Be_True(t *testing.T) {
	expected := true
	userId := "006"

	actualResult := CheckUserInRankByID(userId)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)

	}

}

func Test_CheckUserInRankByID_Input_007_Should_Be_False(t *testing.T) {
	expected := false
	userId := "007"

	actualResult := CheckUserInRankByID(userId)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)

	}

}

func Test_UpdateLevel_Input_UserID_006_Output_Should_be_Platinum(t *testing.T) {
	expected := "Platinum"
	userId := 006

	actualResult := CheckUserID(userId)

	if expected != actualResult {
		t.Errorf("Failed because the expected result is %s but got %s", expected, actualResult)
	}

}
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

}

func Test_UpdatePoint_Input_UserID_006_Level_Platinium_OutPut_Should_Be_800(t *testing.T) {
	userId := "006"
	level := "Platinum"
	point := 800

	expected := updatepoint(userId, level)

	if expected != point {
		t.Errorf("Error! Expect %v but %v", point, expected)
	}
}
func Test_GetLastSixMonthByUserId_Input_006_Should_Be_Transaction_Within_Six_Month(t *testing.T) {
	expectedCountTransactionNumber := 8
	userId := "006"

	actualTransactions := GetLastSixMonthByUserId(userId)

	if len(actualTransactions) != expectedCountTransactionNumber {
		t.Errorf("Error, the transaction should be (%v) but it is (%v)", expectedCountTransactionNumber, len(actualTransactions))
	}
}
