package execute

import (
	"testing"
	"strings"
	"fmt"
	"io/ioutil"
	"os"
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
func Test_generateCommandString(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = SetOptions(options, "")

	actual := generateCommandString(options, option.Param{Src: "src", Dest: "dest"})
	expected := []string{"mv", "--", "src", "dest"}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
	}
}
func Test_generateCommandString_cp(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = SetOptions(options, "C")

	actual := generateCommandString(options, option.Param{Src: "src", Dest: "dest"})
	expected := []string{"cp", "--", "src", "dest"}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
	}
}
func Test_generateCommandString_ln(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = SetOptions(options, "L")

	actual := generateCommandString(options, option.Param{Src: "src", Dest: "dest"})
	expected := []string{"ln", "--", "src", "dest"}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
	}
}
func Test_generateCommandString_ln_s(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = SetOptions(options, "Ls")

	actual := generateCommandString(options, option.Param{Src: "src", Dest: "dest"})
	expected := []string{"ln", "-s", "--", "src", "dest"}
	if Diff(expected, actual) {
		t.Errorf("\ngot : %v\nwant: %v", actual, expected)
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
func Test_ExecuteCommands(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = SetOptions(options, "")

	if err := ioutil.WriteFile("../testdata/case2/hoge1", []byte(""), 0644); err != nil {
		t.Errorf("error: %v", err)
	}
	params := []option.Param{
		{Src: "../testdata/case2/hoge1", Dest: "../testdata/case2/hoge2"},
	}


	err := ExecuteCommands(options, params)
	if err != nil {
		t.Errorf("ExecuteCommands failed: %v", err)
	}

	
	if _, err := os.Stat("../testdata/case2/hoge1"); err == nil {
		//なくなっているはず。
		t.Errorf("file must not exists")
	}
	
	if _, err := os.Stat("../testdata/case2/hoge2"); err != nil {
		//あるはず。
		t.Errorf("file must exists")
	}
	os.Remove("../testdata/case2/hoge2")
}
func Test_ExecuteCommands_overwrite(t *testing.T) {
	options_p := new(option.Option)
	options := *options_p
	options = SetOptions(options, "")

	if err := ioutil.WriteFile("../testdata/case2/hoge1", []byte(""), 0644); err != nil {
		t.Errorf("error: %v", err)
	}
	if err := ioutil.WriteFile("../testdata/case2/hoge2", []byte(""), 0644); err != nil {
		t.Errorf("error: %v", err)
	}
	params := []option.Param{
		{Src: "../testdata/case2/hoge1", Dest: "../testdata/case2/hoge2"},
	}


	err := ExecuteCommands(options, params)
	if err == nil {
		t.Errorf("ExecuteCommands overwrite must error.")
	}

	if _, err := os.Stat("../testdata/case2/hoge1"); err != nil {
		//移動してないから存在するはず
		t.Errorf("file must not exists")
	}
	
	os.Remove("../testdata/case2/hoge1")
	os.Remove("../testdata/case2/hoge2")
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
