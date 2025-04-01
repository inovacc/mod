package main

import (
	"fmt"
	"github.com/inovacc/mod/module"
)

func main() {
	mod := module.NewModule()
	result, err := mod.GetModule("example.com/m", "latest")
	if err != nil {
		panic(err)
	}

	for _, v := range result {
		fmt.Printf("Path: %s, Version: %s\n", v.Path, v.Version)
	}
}
