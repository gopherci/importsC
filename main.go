package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "The first argument should be the import path")
		os.Exit(3)
	}
	importsC, err := ImportsC(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if importsC {
		fmt.Println("Directly or indirectly imports C.")
		os.Exit(0)
	}
	fmt.Println("Does not directly or indirectly import C.")
	os.Exit(1)
}

var checked = make(map[string]bool)

func ImportsC(path string) (bool, error) {
	context := build.Default
	pkg, err := context.Import(path, "", 0)
	if err != nil {
		return false, err
	}

	for _, imprt := range pkg.Imports {
		if imprt == "C" {
			return true, nil
		}
		if ok := checked[imprt]; ok {
			continue
		}
		ok, err := ImportsC(imprt)
		if ok || err != nil {
			return ok, err
		}
		checked[imprt] = true
	}
	return false, nil
}
