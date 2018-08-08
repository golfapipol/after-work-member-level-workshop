package memberlevel

import "testing"

func Test_CheckCountOrder_Input_8_should_Be_True(t *testing.T) {
	expected := true
	order := 8

	actualResult := CheckCountOrder(order)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, order)
	}

}

func Test_CheckCountOrder_Input_8_should_Be_False(t *testing.T) {
	expected := false
	order := 7

	actualResult := CheckCountOrder(order)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, order)
	}

}
