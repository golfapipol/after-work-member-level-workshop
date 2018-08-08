package memberlevel

import "testing"

func Test_UpdateLevel_Input_UserID_006_Output_Should_be_Platinum(t *testing.T) {
	expected := "Platinum"
	userId := 006

	actualResult := CheckUserID(userId)

	if expected != actualResult {
		t.Errorf("Failed becuase the expected result is %s but got %s", expected, actualResult)
	}

}
