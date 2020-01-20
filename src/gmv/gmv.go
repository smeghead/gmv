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
func main() {
	options := option.NewOption()
	flag.Parse()
	params, err := parse(options, flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	for _, p := range params {
		fmt.Printf("%s -- '%s' '%s'\n", "mv", p.Src, p.Dest)
	}
}
