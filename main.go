package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ARC5RF/daisy/types/array"
)

func main() {
	stat, err := os.Stat(".daisyfile")
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("insufficient permission to open .daisyfile")
		} else if os.IsNotExist(err) {
			fmt.Println(".daisyfile does not exist")
		} else {
			fmt.Println("something unexpected happened", err)
		}
		return
	}
	if stat.IsDir() {
		fmt.Println(".daisyfile is a directory")
		return
	}

	data, err := os.ReadFile(".daisyfile")
	if err != nil {
		fmt.Println("could not open .daisyfile")
		return
	}

	var config DaisyFile
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) < 2 {
		config.Usage()
		return
	}

	env := array.Make[string](0)
	if len(os.Args) > 2 {
		env.Push(os.Args[2:]...)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("could not get current working directory")
		return
	}

	if err := config.Do(os.Args[1], wd, env); err != nil {
		fmt.Println(err)
		return
	}
}
