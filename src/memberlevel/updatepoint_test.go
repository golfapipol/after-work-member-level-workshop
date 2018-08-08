package memberlevel

import "testing"

func Test_UpdatePoint_Input_UserID_006_Level_Platinium_OutPut_Should_Be_800(t *testing.T) {
	userId := "006"
	level := "Platinum"
	point := 800

	expected := updatepoint(userId, level)

	if expected != point {
		t.Errorf("Error jaaa. Expect %v but %v", point, expected)
	}
}
