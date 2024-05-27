package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ARC5RF/daisy/types/array"
)

func matches(name string, v fs.DirEntry) bool {
	info, err := v.Info()
	return err == nil && info.Mode()&0111 != 0 && v.Name() == name
}

func only_matches(name, candidate string) []fs.DirEntry {
	entries, err := os.ReadDir(candidate)
	if err != nil {
		panic(err)
	}
	output := []fs.DirEntry{}
	for _, entry := range entries {
		if matches(name, entry) {
			output = append(output, entry)
		}
	}
	return output
}

func Which(name string) []string {
	output := []string{}
	p := os.Getenv("PATH")

	candidates := strings.Split(p, string(os.PathListSeparator))
	for _, candidate := range candidates {
		if s, err := os.Stat(candidate); os.IsNotExist(err) || !s.IsDir() {
			continue
		}
		for _, entry := range only_matches(name, candidate) {
			output = append(output, filepath.Join(candidate, entry.Name()))
		}
	}

	return output
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"

var l sync.Mutex
var pp = ""
var pf = ""

func print_with_prefix(prefix, location, message, f, c string) {
	var current_prefix = prefix
	l.Lock()
	if pp != prefix || pf != f {
		a := fmt.Sprintf("%s%s%s", f, prefix, Reset)
		b := fmt.Sprintf("%s%s%s", Magenta, location, Reset)
		current_prefix = fmt.Sprintf("%s %s %s\n", a, c, b)
		pp = prefix
		pf = f
	} else {
		current_prefix = ""
	}
	fmt.Printf("%s%s", current_prefix, message)
	l.Unlock()
}

func output(prefix, location string, reader io.ReadCloser, f, c string) error {
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			print_with_prefix(prefix, location, string(buf[:num]), f, c)
		}
	}
}

func fix_env(env, env_overrides *array.Of[string]) []string {
	lut := map[string]string{}
	for _, v := range os.Environ() {
		if idx := strings.Index(v, "="); idx > 0 {
			ek := v[:idx]
			ev := v[idx+1:]
			lut[ek] = ev
		}
	}
	if env_overrides != nil {
		for _, v := range *env_overrides {
			if idx := strings.Index(v, "="); idx > 0 {
				ek := v[:idx]
				ev := v[idx+1:]
				lut[ek] = ev
			}
		}
	}
	if env != nil {
		for _, v := range *env {
			if idx := strings.Index(v, "="); idx > 0 {
				ek := v[:idx]
				ev := v[idx+1:]
				lut[ek] = ev
			}
		}
	}
	output := []string{}
	for k, v := range lut {
		output = append(output, k+"="+v)
	}
	return output
}

func Execute(filename, prefix, location string, env, env_overrides, args *array.Of[string]) error {
	cmd := exec.Command(filename, (*args)...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if env.Length() > 0 || env_overrides.Length() > 0 {
		cmd.Env = fix_env(env, env_overrides)
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error executing command: %s......\n", err.Error())
		return err
	}

	go output(prefix, location, stdout, Green, filepath.Base(filename))
	go output(prefix, location, stderr, Red, filepath.Base(filename))

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for command execution: %s......\n", err.Error())
		return err
	}

	return nil
}
