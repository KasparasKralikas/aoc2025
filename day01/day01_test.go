package main

import "testing"

func TestSolveSample(t *testing.T) {
	firstPassword, secondPassword, err := solve("sample.txt")
	if err != nil {
		t.Fatalf("solve(sample.txt) returned error: %v", err)
	}

	const want1 = 3
	if firstPassword != want1 {
		t.Fatalf("First password = %d, want %d", firstPassword, want1)
	}

	const want2 = 6
	if secondPassword != want2 {
		t.Fatalf("Second password = %d, want %d", secondPassword, want2)
	}
}
