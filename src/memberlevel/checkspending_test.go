package memberlevel

import (
	"testing"
)

func Test_CheckSpending_Input_1200_Should_Be_True(t *testing.T) {
	spending := 1200
	expected := true

	actual := CheckSpending(spending)

	if expected != actual {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}

func Test_CheckSpending_Input_2000_Should_Be_False(t *testing.T) {
	spending := 2000
	expected := true

	actual := CheckSpending(spending)

	if expected != actual {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}

func Test_CheckSpending_Input_1001_Should_Be_False(t *testing.T) {
	spending := 1001
	expected := true

	actual := CheckSpending(spending)

	if expected != actual {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}
