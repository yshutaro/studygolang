package main

import "testing"

func TestReverse2(t *testing.T) {
	if reverse2(123) != 321 {
		t.Error(`false`)
	}
	if reverse2(-123) != -321 {
		t.Error(`false`)
	}
	if reverse2(120) != 21 {
		t.Error(`false`)
	}
	if reverse2(1534236469) != 0 {
		t.Error(`false`)
	}
}
