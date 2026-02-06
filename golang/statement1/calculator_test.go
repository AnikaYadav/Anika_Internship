package main

import "testing"

func TestCalculateAdd(t *testing.T) {
	result, err := calculate(2, 3, "+")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %v", result)
	}
}

func TestCalculateSubtract(t *testing.T) {
	result, err := calculate(5, 3, "-")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("expected 2, got %v", result)
	}
}

func TestCalculateMultiply(t *testing.T) {
	result, err := calculate(4, 5, "*")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 20 {
		t.Errorf("expected 20, got %v", result)
	}
}

func TestCalculateDivide(t *testing.T) {
	result, err := calculate(10, 2, "/")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %v", result)
	}
}

func TestCalculateDivideByZero(t *testing.T) {
	_, err := calculate(10, 0, "/")
	if err == nil {
		t.Errorf("expected error when dividing by zero")
	}
}

func TestCalculateInvalidOperator(t *testing.T) {
	_, err := calculate(10, 5, "%")
	if err == nil {
		t.Errorf("expected error for invalid operator")
	}
}
