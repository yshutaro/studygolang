package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{"flower", "flow", "flight"}) != "fl" {
		t.Error(`false`)
	}
	if longestCommonPrefix([]string{"dog", "racecar", "car"}) != "" {
		t.Error(`false`)
	}
	if longestCommonPrefix([]string{"a"}) != "a" {
		t.Error(`false`)
	}
	if longestCommonPrefix([]string{}) != "" {
		t.Error(`false`)
	}
}
