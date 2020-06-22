package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(121) {
		t.Error(`false`)
	}
	if isPalindrome(-121) {
		t.Error(`false`)
	}
	if isPalindrome(10) {
		t.Error(`false`)
	}
	if !isPalindrome(9646336469) {
		t.Error(`false`)
	}
}
