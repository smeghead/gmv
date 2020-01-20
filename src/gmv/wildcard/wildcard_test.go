package wildcard

import (
	"strings"
	"testing"
	"fmt"
	"gmv/option"
)

func setOptions(options option.Option, param string) option.Option {
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
func Test_parse_no_match(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	fmt.Println(options)
	options = setOptions(options, "")
	ret, err := Parse(options, "testdata/(case1/(*).c)")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	actual := ret
	expected := []PathElement{
		{charType: Literal, content: "testdata/", match: "", referenceNumbers: []int{}},
		{charType: Literal, content: "case1/", match: "", referenceNumbers: []int{1}},
		{charType: Star, content: "*", match: "", referenceNumbers: []int{1, 2}},
		{charType: Literal, content: ".c", match: "", referenceNumbers: []int{1}},
	}
	if len(actual) != len(expected) {
		t.Errorf("\ngot : %v\nwant: %v", len(actual), len(expected))
	}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
	}
}
func Test_parse_simple_match(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = setOptions(options, "")
	actual, err := Parse(options, "testdata/case1/(*)")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	expected := []PathElement{
		{charType: Literal, content: "testdata/case1/", match: "", referenceNumbers: []int{}},
		{charType: Star, content: "*", match: "", referenceNumbers: []int{1}},
	}
	if len(actual) != len(expected) {
		t.Errorf("\ngot : %v\nwant: %v", len(actual), len(expected))
	}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
	}
}
func Test_parse_w_match(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = setOptions(options, "w")
	actual, err := Parse(options, "testdata/case1/*")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	expected := []PathElement{
		{charType: Literal, content: "testdata/case1/", match: "", referenceNumbers: []int{}},
		{charType: Star, content: "*", match: "", referenceNumbers: []int{1}},
	}
	if len(actual) != len(expected) {
		t.Errorf("\ngot : %v\nwant: %v", len(actual), len(expected))
	}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
	}
}
