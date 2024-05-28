package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ARC5RF/daisy/types/array"
)

type DaisyLocation string

func (location *DaisyLocation) UnmarshalJSON(data []byte) error {
	var temp any
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	switch v := temp.(type) {
	case string:
		*location = DaisyLocation(v)
		l := string(*location)
		if strings.HasPrefix(l, ".") {
			return nil
		}
		if strings.HasPrefix(l, "..") {
			return nil
		}
		if strings.HasPrefix(l, "daisy.") {
			return nil
		}
		return errors.New("invalid prefix for field \"in\", expected a relative path or daisy.<builtin directive>")
	case map[string]any:
		return errors.New("field \"in\" does not support object values yet")
	default:
		return errors.New("field \"in\" does not support " + reflect.TypeOf(temp).String())
	}
}

func (location DaisyLocation) Candidates() *array.Of[string] {
	output := array.Make[string](0)

	l := string(location)
	if strings.HasPrefix(l, "daisy.") {
		return path_candidates_for_directive(l, output)
	}

	output.Push(string(location))

	return output
}

type DaisyCommand struct {
	Command  string            `json:"command"`
	Prefix   string            `json:"prefix"`
	Detached bool              `json:"detached"`
	Args     *array.Of[string] `json:"args"`
	Env      *array.Of[string] `json:"env"`
}

func env_extract(key string, env, env_overrides *array.Of[string]) string {

	for i := 0; i < env_overrides.Length(); i++ {
		v := env_overrides.At(i)
		if v == nil {
			continue
		}
		if strings.HasPrefix(*v, key+"=") {
			return (*v)[len(key)+1:]
		}
	}

	for i := 0; i < env.Length(); i++ {
		v := env.At(i)
		if v == nil {
			continue
		}
		if strings.HasPrefix(*v, key+"=") {
			return (*v)[len(key)+1:]
		}
	}

	os_env := os.Environ()
	for i := 0; i < len(os_env); i++ {
		v := os_env[i]
		if strings.HasPrefix(v, key+"=") {
			return (v)[len(key)+1:]
		}
	}

	return ""
}

func output_ext(env, env_overrides *array.Of[string]) string {
	goos := env_extract("GOOS", env, env_overrides)
	goarch := env_extract("GOARCH", env, env_overrides)

	if goos == "windows" {
		return ".exe"
	}

	if goos == "js" && goarch == "wasm" {
		return ".wasm"
	}

	return ""
}

func (command DaisyCommand) FixArgs(wd, location string, env_overrides *array.Of[string]) {
	for i := 0; i < command.Args.Length(); i++ {
		v := command.Args.At(i)
		*v = strings.ReplaceAll(*v, "$DAISY_ROOT", wd)
		*v = strings.ReplaceAll(*v, "$DAISY_WD", location)
		*v = strings.ReplaceAll(*v, "$DAISY_OS_EXT", output_ext(command.Env, env_overrides))
	}
}

func run_internal_execute(command DaisyCommand, p, location string, env_overrides *array.Of[string]) error {
	bins := Which(command.Command)
	if len(bins) == 0 {
		return errors.New("could not find " + command.Command)
	}

	e := ""
	if env_overrides.Length() > 0 {
		e = strings.Join(*env_overrides, " ") + " "
	}

	m := fmt.Sprintln(e+filepath.Base(bins[0]), strings.Join((*command.Args), " "))
	print_with_prefix(p, location, m, Green, filepath.Base(bins[0]))

	if command.Detached {
		go Execute(bins[0], p, location, command.Env, env_overrides, command.Args)
		return nil
	}

	return Execute(bins[0], p, location, command.Env, env_overrides, command.Args)
}

func run_internal_directive(command DaisyCommand, p, location string, env_overrides *array.Of[string]) error {
	switch command.Command {
	case "daisy.move.force":
		return move_force(command, p, location, env_overrides)
	}

	return nil
}

func run_internal(command DaisyCommand, pre, location string, env_overrides *array.Of[string]) error {
	p := pre
	if len(command.Prefix) > 0 {
		p = command.Prefix
	}

	if strings.HasPrefix(command.Command, "daisy.") {
		return run_internal_directive(command, pre, location, env_overrides)
	}
	return run_internal_execute(command, p, location, env_overrides)
}

func (command DaisyCommand) Run(wd, location, pre string, env_overrides *array.Of[string]) error {
	command.FixArgs(wd, location, env_overrides)
	if err := os.Chdir(location); err != nil {
		return err
	}

	return run_internal(command, pre, location, env_overrides)
}

type DaisyCommands []DaisyCommand

func (commands DaisyCommands) Run(wd, location, pre string, env_overrides *array.Of[string]) error {
	for _, command := range commands {
		if err := command.Run(wd, location, pre, env_overrides); err != nil {
			return err
		}
	}
	return nil
}

type DaisyTask struct {
	In DaisyLocation `json:"in"`
	Do DaisyCommands `json:"do"`
}

func (task DaisyTask) Run(wd, pre string, env_overrides *array.Of[string]) error {
	for _, location := range *task.In.Candidates() {
		l := filepath.Join(wd, location)
		if err := task.Do.Run(wd, l, pre, env_overrides); err != nil {
			return err
		}
	}
	return nil
}

type DaisyTaskList []DaisyTask

func (list DaisyTaskList) Run(wd, pre string, env_overrides *array.Of[string]) error {
	for _, v := range list {
		if err := v.Run(wd, pre, env_overrides); err != nil {
			return err
		}
	}
	return nil
}

type DaisyTasks map[string]DaisyTaskList

type DaisyFile struct {
	Version  string     `json:"version"`
	Commands DaisyTasks `json:"commands"`
}

func (file DaisyFile) Usage() {
	fmt.Println("usage:\n\tdaisy command")
	fmt.Println("commands")
	for k := range file.Commands {
		fmt.Println("\t" + k)

	}
}

func (file DaisyFile) Do(task, wd string, env_overrides *array.Of[string]) error {
	cmd, has_cmd := file.Commands[task]
	if !has_cmd {
		return errors.New("command " + task + " does not exist")
	}

	return cmd.Run(wd, task, env_overrides)
}
