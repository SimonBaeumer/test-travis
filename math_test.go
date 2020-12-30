package main

import "testing"

func TestAdd(t *testing.T) {
	r := Add(2, 2)
	if r != 3 {
		t.Error("Result was not 3.")
	}
}
