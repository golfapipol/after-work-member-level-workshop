package memberlevel

import "testing"

func Test_CheckCountOrder_Input_8_should_Be_True(t *testing.T) {
	expected := true
	orderId := 8

	actualResult := CheckCountOrder(orderId)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, orderId)
	}

}

func Test_CheckCountOrder_Input_8_should_Be_False(t *testing.T) {
	expected := false
	orderId := 7

	actualResult := CheckCountOrder(orderId)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, orderId)
	}

}
