package mymath

import "testing"

// go test -v
func TestSum(t *testing.T) {

	// test for basic operation
	sumResult := Sum(2,5)
	if sumResult != 7 {
		t.Errorf("Sum(2,5) failed, expected %d, got %d", 7, sumResult)
	} else {
		t.Logf("Sum(2,5) succeed, expected %d, got %d", 7, 7)
	}

	abstractResult := Abstract(5,3)
	if abstractResult != 2 {
		t.Errorf("Abstract(5,3) failed, expected %d, got %d", 2, abstractResult)
	} else {
		t.Logf("Abstract(5,3) succeed, expected %d, got %d", 2, 2)
	}

	multiplyResult := Multiply(9,6)
	if multiplyResult != 54 {
		t.Errorf("Multiply(9,6) failed, expected %d, got %d", 54, multiplyResult)
	} else {
		t.Logf("Multiply(9,6) succeed, expected %d, got %d", 54, 54)
	}

	divideResult := Divide(15,5)
	if divideResult != 3 {
		t.Errorf("Divide(15,5) failed, expected %d, got %d", 3, divideResult)
	} else {
		t.Logf("Divide(15,5) succeed, expected %d, got %d", 3, 3)
	}
}