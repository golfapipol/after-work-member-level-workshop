package memberlevel

import (
	"testing"
)

func Test_GetLastSixMonthByUserId_Input_006_Should_Be_Transaction_Within_Six_Month(t *testing.T) {
	expectedCountTransactionNumber := 8
	userId := "006"

	actualTransactions := GetLastSixMonthByUserId(userId)

	if len(actualTransactions) != expectedCountTransactionNumber {
		t.Errorf("Error, the transaction should be (%v) but it is (%v)", expectedCountTransactionNumber, len(actualTransactions))
	}
}
