package execute

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
	"gmv/option"
)

func quoted(s string) string {
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return fmt.Sprintf("\"%s\"", s)
}
func generateCommandString(options option.Option, param option.Param) []string {
	args := []string{}
	program := ""
	switch {
	case *options.Opt_C:
		program = "cp"
	case *options.Opt_L:
		program = "ln"
	default:
		program = "mv"
	}
	args = append(args, program)
	//option
	args = append(args, "--")
	args = append(args, quoted(param.Src), quoted(param.Dest))
	return args
}

func checkOverride(params []option.Param) error {
	paths := make(map[string]int)
	for _, p := range params {
		paths[p.Src] += 1
		paths[p.Dest] += 1
	}
	for path, count := range paths {
		if count > 1 {
			return fmt.Errorf("duplicate paths. [%s]", path)
		}
	}
	return nil
}
func ExecuteCommands(options option.Option, params []option.Param) {
	if *options.Opt_n {
		for _, p := range params {
			fmt.Printf("%s -- '%s' '%s'\n", "mv", p.Src, p.Dest)
		}
		return
	}

	if err := checkOverride(params); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v.\n", err)
		return
	}
	//ファイル移動
	for _, p := range params {
		if err := exec.Command("mv", "--", p.Src, p.Dest).Run(); err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
			return
		}
	}
}
