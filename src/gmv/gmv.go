package main

import (
	"os"
	"flag"
	"fmt"
	"github.com/mattn/go-zglob"
	"gmv/option"
	"gmv/wildcard"
)

type Param struct {
	Src string
	Dest string
}

func checkOverride(params []Param) error {
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


func parse(options option.Option, src, dest string) ([]Param, error) {
	elements, err := wildcard.Parse(options, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return nil, err
	}
	matches, err := zglob.Glob(wildcard.GetSearchPath(elements))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return nil, err
	}
	params := []Param{}
	for _ , path := range matches {
		destPath, err := wildcard.GetDestPath(options, elements, path, dest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
			return nil, err
		}
		params = append(params, Param{Src: path, Dest: destPath})
	}
	return params, nil
}
func executeCommands(options option.Option, params []Param) {
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
	//TODO ファイル移動は実行していない。
}
func main() {
	options := option.NewOption()
	flag.Parse()
	params, err := parse(options, flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	if len(params) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: no target files\n")
		return
	}
	
	executeCommands(options, params)
}
