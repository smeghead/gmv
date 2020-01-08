package main

import (
	"flag"
	"fmt"
)

func fun() string {
	return "aa"
}
func main() {
	flag.Parse()
	var ret = fun()
	fmt.Println(ret)
}
