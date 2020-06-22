package main

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("I") != 1 {
		t.Error(`false`)
	}
	if romanToInt("III") != 3 {
		t.Error(`false`)
	}
	if romanToInt("IV") != 4 {
		t.Error(`false`)
	}
	if romanToInt("IX") != 9 {
		t.Error(`false`)
	}
	if romanToInt("LVIII") != 58 {
		t.Error(`false`)
	}
	if romanToInt("MCMXCIV") != 1994 {
		t.Error(`false`)
	}
}
