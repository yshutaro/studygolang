package main

import "testing"

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.Error(`false`)
	}
	if reverse(-123) != -321 {
		t.Error(`false`)
	}
	if reverse(120) != 21 {
		t.Error(`false`)
	}
	if reverse(1534236469) != 0 {
		t.Error(`false`)
	}
}
