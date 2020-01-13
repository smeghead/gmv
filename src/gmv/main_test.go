package main

import (
	"testing"
	"reflect"
)

func Test_parse_no_match(t *testing.T) {
	options := new(Option)
	ret, err := parse(options, "testdata/case1/*.c", "testdata/case/*.x")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	actual := ret
	expected := []Param{}
	if len(actual) != len(expected) {
		t.Errorf("got: %v\nwant: %v", len(actual), len(expected))
	}
//	if reflect.DeepEqual(expected, actual) {
//		t.Errorf("got: %v\nwant: %v", actual, expected)
//	}
}
func Test_parse_simple_match(t *testing.T) {
	options := new(Option)
	actual, err := parse(options, "testdata/case1/*", "testdata/case1/*.x")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	expected := []Param{
		{Src: "testdata/case1/01.txt", Dest: "testdata/case1/01.txt.x"},
		{Src: "testdata/case1/02.txt", Dest: "testdata/case1/02.txt.x"},
		{Src: "testdata/case1/03.txt", Dest: "testdata/case1/03.txt.x"},
	}
	if len(actual) != len(expected) {
		t.Errorf("got: %v\nwant: %v", len(actual), len(expected))
	}
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
