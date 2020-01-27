package execute

import (
	"os"
	"io"
	"os/exec"
	"strings"
	"fmt"
	"io/ioutil"
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
	if (*options.Opt_L && *options.Opt_s) {
		args = append(args, "-s")
	}
	args = append(args, "--")
	args = append(args, param.Src, param.Dest)
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
func ExecuteCommands(options option.Option, params []option.Param) error {
	if *options.Opt_n {
		for _, p := range params {
			command := generateCommandString(options, p)
			fmt.Printf("%s\n", strings.Join(command, " "))
		}
		return nil
	}

	if err := checkOverride(params); err != nil {
		return err
	}
	//ファイル移動
	for _, p := range params {
		if _, err := os.Stat(p.Dest); err == nil {
			// 移動後のファイルが存在する場合
			return fmt.Errorf("target file exists. %s", p.Dest)
		}

		commandStrings := generateCommandString(options, p)
		command := exec.Command(commandStrings[0], commandStrings[1:]...)
		if *options.Opt_i || *options.Opt_v {
			fmt.Printf("%s\n", command.String())
		}
		if *options.Opt_i {
			fmt.Fprintf(os.Stderr, "Execute? (y/n)")
			input := ""
			if _, err := fmt.Scanf("%s", &input); err != nil {
				return err
			}
			if len(input) == 0 {
				continue
			}
			input = strings.ToLower(input)
			if input[0] != 'y' {
				continue
			}
		}
		stdout, err := command.StdoutPipe()
		if err != nil {
			return err
		}
		stderr, err := command.StderrPipe()
		if err != nil {
			return err
		}

		if err := command.Start(); err != nil {
			return err
		}

		outputStream := func(s io.ReadCloser, w io.Writer){
			slurp, _ := ioutil.ReadAll(s)
			if len(slurp) > 0 {
				fmt.Fprintf(w, "%s\n", slurp)
			}
		}
		outputStream(stderr, os.Stderr)
		outputStream(stdout, os.Stdout)

		if err := command.Wait(); err != nil {
			return err
		}
	}
	return nil
}
