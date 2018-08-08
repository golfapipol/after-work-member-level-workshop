package memberlevel

import "testing"

func Test_CheckUserInRankByID_Input_006_Should_Be_True(t *testing.T){
	expected := true
	userId := "006"

	actualResult := CheckUserInRankByID(userId)

	if expected != actualResult{
		t.Errorf("Expected %v but got it %v", expected, actualResult)

	}


}

func Test_CheckUserInRankByID_Input_007_Should_Be_False(t *testing.T){
	expected := false
	userId := "007"

	actualResult := CheckUserInRankByID(userId)

	if expected != actualResult{
		t.Errorf("Expected %v but got it %v", expected, actualResult)

	}


}