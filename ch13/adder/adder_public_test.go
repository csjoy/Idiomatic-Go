package adder_test

import (
	"main/adder"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Errorf("incorrect result: expected 5, got %v\n", result)
	}

}
