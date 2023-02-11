package table_test

import (
	"main/table"
	"testing"
)

// func TestDoMath(t *testing.T) {
// 	result, err := table.DoMath(2, 2, "+")
// 	if err != nil {
// 		t.Error("Should have been nil error, got", err)
// 	}
// 	if result != 4 {
// 		t.Error("Should have been 4, got", result)
// 	}

// 	result2, err2 := table.DoMath(2, 2, "-")
// 	if err2 != nil {
// 		t.Error("Should have been nil error, got", err2)
// 	}
// 	if result2 != 0 {
// 		t.Error("Should have been 0, got", result2)
// 	}
// 	// and so on
// }

func TestDoMathTable(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, "division by zero"},
		{"bad_op", 2, 2, "?", 0, "unknown operator ?"},
		{"another_mult", 2, 3, "*", 6, ""},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := table.DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
