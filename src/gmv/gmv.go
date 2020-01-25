package main

import (
	"os"
	"flag"
	"fmt"
	"github.com/mattn/go-zglob"
	"gmv/option"
	"gmv/wildcard"
	"gmv/execute"
)




func parse(options option.Option, src, dest string) ([]option.Param, error) {
	elements, err := wildcard.Parse(options, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		return nil, err
	}
	matches, err := zglob.Glob(wildcard.GetSearchPath(elements))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		return nil, err
	}
	params := []option.Param{}
	for _ , path := range matches {
		destPath, err := wildcard.GetDestPath(options, elements, path, dest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
			return nil, err
		}
		params = append(params, option.Param{Src: path, Dest: destPath})
	}
	return params, nil
}
func main() {
	options := option.NewOption()
	flag.Parse()

	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(-1)
		return
	}

	params, err := parse(options, flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(-1)
		return
	}
	if len(params) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: no target files\n")
		os.Exit(-1)
		return
	}
	
	err = execute.ExecuteCommands(options, params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: no target files\n")
		os.Exit(-1)
		return
	}
}
