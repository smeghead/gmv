package main

import (
	"testing"
)

func Test_func(t *testing.T) {
	var ret = fun()
	actual := ret
	expected := "aaa"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
