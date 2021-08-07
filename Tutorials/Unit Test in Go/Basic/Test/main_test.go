package main

import "testing"

func TestBasic(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculate(t *testing.T) {

	// Declare the parameter and returns of the Function Calculate()
	var tests = []struct {
		input  int // parameter
		expect int // returns
	}{
		// Preset test data {parameter,return}
		{99, 101},
		{2, 5},
		{4, 6},
	}
	// Run the test for Function Calculate
	for _, test := range tests {
		// Checks the function if the return is what we expect, if not it will promt t.Errorf
		if output := Calculate(test.input); output != test.expect {
			// whenever theres an error means the test will result to fail
			t.Errorf("Test Failed: %d inputted, %d expected, %d recieved ", test.input, test.expect, output)
		}
	}
}
