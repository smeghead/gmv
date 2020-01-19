package gmv

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
	fmt.Println("Src: ", src)
	fmt.Println("Dest: ", dest)
	elements, err := wildcard.Parse(options, src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return nil, err
	}
	fmt.Println(elements)
	matches, err := zglob.Glob(wildcard.GetSearchPath(elements))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return nil, err
	}
	fmt.Println("Count:", len(matches))
	params := []Param{}
	for _ , path := range matches {
		fmt.Println(path)
		destPath, err := wildcard.GetDestPath(options, elements, path, dest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
			return nil, err
		}
		fmt.Println(destPath)
		params = append(params, Param{Src: path, Dest: destPath})
	}
	return params, nil
}
func main() {
	options := option.NewOption()
	flag.Parse()
	fmt.Println(options.Opt_o)
	params, err := parse(options, flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	fmt.Println(params)
}
