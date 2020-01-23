package execute

import (
	"testing"
	"strings"
	"fmt"
	"gmv/option"
)

func SetOptions(options option.Option, param string) option.Option {
	t := true
	f := false
	options.Opt_W = &t
	options.Opt_w = &f
	options.Opt_f = &f
	options.Opt_i = &f
	options.Opt_n = &f
	options.Opt_q = &f
	options.Opt_Q = &f
	options.Opt_s = &f
	options.Opt_v = &f
	options.Opt_w = &f
	options.Opt_W = &f
	options.Opt_C = &f
	options.Opt_L = &f
	options.Opt_M = &f
//	Opt_o *string
//	Opt_p *string
//	Opt_P *string

	params := strings.Split(param, "")
	for _, p := range params {
		switch p {
		case "f":  options.Opt_f = &t; break
		case "i":  options.Opt_i = &t; break
		case "n":  options.Opt_n = &t; break
		case "q":  options.Opt_q = &t; break
		case "Q":  options.Opt_Q = &t; break
		case "s":  options.Opt_s = &t; break
		case "v":  options.Opt_v = &t; break
		case "w":  options.Opt_w = &t; break
		case "W":  options.Opt_W = &t; break
		case "C":  options.Opt_C = &t; break
		case "L":  options.Opt_L = &t; break
		case "M":  options.Opt_M = &t; break

		}
	}
	return options
}
func Diff(expected, actual interface{}) bool {
	return fmt.Sprintf("%#v", actual) != fmt.Sprintf("%#v", expected)
}
type testCase struct {
	org string
	expected string
}
func Test_quoted(t *testing.T) {
	cases := []testCase{
		{"testdata/case1/01.txt", "\"testdata/case1/01.txt\""},
		{"testdata/case1/\".txt", "\"testdata/case1/\\\".txt\""},
		{"testdata/case1/0 .txt", "\"testdata/case1/0 .txt\""},
	}
	for _, c := range cases {
		actual := quoted(c.org)
		if Diff(c.expected, actual) {
			t.Errorf("\ngot : %v\nwant: %v", actual, c.expected)
		}
	}
}
func Test_checkOverride(t *testing.T) {
	params := []option.Param{
		{Src: "hoge1", Dest: "hoge2"},
		{Src: "hoge3", Dest: "hoge4"},
	}
	actual := checkOverride(params)

	if actual != nil {
		t.Errorf("got: %v", actual)
	}
}
func Test_checkOverride_duplicates(t *testing.T) {
	params := []option.Param{
		{Src: "hoge1", Dest: "hoge2"},
		{Src: "hoge2", Dest: "hoge4"},
	}
	actual := checkOverride(params)

	if actual == nil {
		t.Errorf("must rase error: %v", actual)
	}
}
//func Test_Parse_simple_match(t *testing.T) {
//	options_p := new(option.Option)
//	options := *options_p
//	options = SetOptions(options, "")
//	actual, err := Parse(options, "testdata/case1/(*)")
//	if err != nil {
//		t.Errorf("err: %v", err)
//	}
//	expected := []PathElement{
//		{charType: Literal, content: "testdata/case1/", match: "", referenceNumbers: []int{}},
//		{charType: Star, content: "*", match: "", referenceNumbers: []int{1}},
//	}
//	if len(actual) != len(expected) {
//		t.Errorf("\ngot : %v\nwant: %v", len(actual), len(expected))
//	}
//	if Diff(expected, actual) {
//		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
//	}
//}
