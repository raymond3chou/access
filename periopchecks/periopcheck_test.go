package periopcheck

import "testing"

func TestCheckValidNumber(t *testing.T) {
	if CheckValidNumber(0, 2, "3") {
		t.Errorf("Function failed\n")
	}
	if !CheckValidNumber(1, 2, "2") {
		t.Error("Upper bound failed 2\n")
	}
	if !CheckValidNumber(-1, 4, "-1") {
		t.Error("Negative lower bound failed -1\n")
	}
	if !CheckValidNumber(1, 5, "3") {
		t.Error("Bounds failed 3")
	}
	if CheckValidNumber(1, 2, "Jim") {
		t.Error("String conversion failed")
	}
}
