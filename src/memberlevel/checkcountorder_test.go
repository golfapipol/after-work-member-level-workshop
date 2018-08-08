package memberlevel

import "testing"

func Test_CheckCountOrder_Input_8_should_Be_True(t *testing.T) {
	expected := true
	order := 8
	month := 6

	actualResult := CheckCountOrder(order,month)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)
	}

}

func Test_CheckCountOrder_Input_8_should_Be_False(t *testing.T) {
	expected := false
	order := 7
	month := 7

	actualResult := CheckCountOrder(order,month)

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


