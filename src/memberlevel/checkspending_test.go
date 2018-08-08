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

func Test_CheckSpending_Input_800_Should_Be_False(t *testing.T) {
	spending := 800
	expected := false

	actual := CheckSpending(spending)

	if expected != actual {
		t.Errorf("Expected is %v but got %v", expected, actual)
	}
}
