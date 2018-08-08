package memberlevel_test

import (
	"testing"
)

func Test_GetFreePoint_Input_Platinum_Should_Be_300(t *testing.T) {
	expected := 300
	level := "Platinum"

	actualResult := CheckLevel(level)

	if expected != actualResult {
		t.Errorf("expected %d but got it %d", expected, actualResult)
	}
}

func Test_GetFreePoint_Input_Gold_Should_Be_100(t *testing.T) {
	expected := 100
	level := "Gold"

	actualResult := CheckLevel(level)

	if expected != actualResult {
		t.Errorf("expected %d but got it %d", expected, actualResult)
	}
}

func Test_GetFreePoint_Input_A_Should_Be_0(t *testing.T) {
	expected := 0
	level := "A"

	actualResult := CheckLevel(level)

	if expected != actualResult {
		t.Errorf("expected %d but got it %d", expected, actualResult)
	}
}
