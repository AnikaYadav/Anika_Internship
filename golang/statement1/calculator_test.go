package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name        string
		op1, op2    float64
		op          rune
		expected    float64
		expectedErr bool
	}{
		{"Addition", 3.0, 5.6, '+', 8.60, false},
		{"Subtraction", 4, 2, '-', 2.00, false},
		{"Multiplication", -7, -2, '*', 14.00, false},
		{"Division", 7, 3.5, '/', 2.00, false},
		{"DivisionByZero", 6, 0, '/', 0, true},
		{"InvalidOperator", 8, 8, '_', 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := calculate(tt.op1, tt.op2, tt.op)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("%s: expected error but got none", tt.name)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: unexpected error: %v", tt.name, err)
				return
			}

			if output != tt.expected {
				t.Errorf("%s: output = %v, expected = %v", tt.name, output, tt.expected)
			}
		})
	}
}
