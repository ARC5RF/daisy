package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ARC5RF/daisy/types/array"
)

const DAISY_GLOB_DIR = "daisy.glob.dir/"
const DAISY_REGEXP_DIR = "daisy.regexp.dir/"

func path_candidates_for_directive(p string, output *array.Of[string]) *array.Of[string] {
	wd, err := os.Getwd()
	if err != nil {
		return output
	}
	if strings.HasPrefix(p, DAISY_REGEXP_DIR) {
		pattern := p[len(DAISY_REGEXP_DIR):]
		pat := regexp.MustCompile(pattern)
		filepath.Walk(wd, func(path string, info fs.FileInfo, err error) error {
			if pat.MatchString(path) {
				output.Push(filepath.Dir(path))
			}

			return nil
		})
		return output
	}
	if strings.HasPrefix(p, DAISY_GLOB_DIR) {
		pattern := p[len(DAISY_GLOB_DIR):]
		hits, err := Glob("./" + pattern)
		if err != nil {
			fmt.Println(err)
			return output
		}

		output.Push(*array.Map(hits, func(v *string, k int, input *array.Of[string]) string {
			return filepath.Dir(*v)
		})...)
	}

	return output
}

func move_replace(command DaisyCommand, p, location string, _ *array.Of[string]) error {
	if command.Args.Length() < 2 {
		return errors.New("daisy.move.force expected two arguments")
	}

	print_with_prefix(p, location, fmt.Sprintln("moving", fmt.Sprint(strings.Join(*command.Args, " to "))), Green, command.Command)

	a := *command.Args.At(0)
	as, err := os.Stat(a)
	if err != nil {
		return err
	}

	b := *command.Args.At(1)
	_, err = os.Stat(b)

	d := filepath.Dir(b)

	if os.IsPermission(err) {
		return err
	}

	if os.IsNotExist(err) {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		print_with_prefix(p, location, b+" exists, replacing\n", Yellow, command.Command)
		if err := os.RemoveAll(b); err != nil {
			return err
		}
		if as.IsDir() {
			os.MkdirAll(b, os.ModePerm)
		} else {
			os.MkdirAll(d, os.ModePerm)
		}
	}

	return os.Rename(a, b)
}
