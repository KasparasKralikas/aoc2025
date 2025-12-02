package main

import "testing"

func TestSolveSample(t *testing.T) {
	sumOfInvalidIds1, sumOfInvalidIds2, err := solve("sample.txt")
	if err != nil {
		t.Fatalf("solve(sample.txt) returned error: %v", err)
	}

	const want1 = 1227775554
	if sumOfInvalidIds1 != want1 {
		t.Fatalf("Sum of invalid IDs = %d, want %d", sumOfInvalidIds1, want1)
	}

	const want2 = 4174379265
	if sumOfInvalidIds2 != want2 {
		t.Fatalf("Sum of invalid IDs = %d, want %d", sumOfInvalidIds1, want2)
	}
}
