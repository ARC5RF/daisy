package main

// "borrowed" from https://github.com/yargevad/filepathx/blob/master/filepathx.go

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ARC5RF/daisy/types/array"
)

// Globs represents one filepath glob, with its elements joined by "**".
// type Globs []string

// Glob adds double-star support to the core path/filepath Glob function.
// It's useful when your globs might have double-stars, but you're not sure.
func Glob(pattern string) (*array.Of[string], error) {
	if !strings.Contains(pattern, "**") {
		// passthru to core package if no double-star
		r, err := filepath.Glob(pattern)
		return array.Wrap(&r), err
	}
	return Expand(strings.Split(pattern, "**"))
}

// Expand finds matches for the provided Globs.
func Expand(globs []string) (*array.Of[string], error) {
	var matches = array.Wrap(&[]string{""}) // accumulate here
	for _, glob := range globs {
		var hits = array.Make[string](0)
		var hitMap = map[string]bool{}
		for _, match := range *matches {
			paths, err := filepath.Glob(match + glob)
			if err != nil {
				return nil, err
			}
			for _, path := range paths {
				err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					// save deduped match from current iteration
					if _, ok := hitMap[path]; !ok {
						hits.Push(path)
						hitMap[path] = true
					}
					return nil
				})
				if err != nil {
					return nil, err
				}
			}
		}
		matches = hits
	}

	// fix up return value for nil input
	if globs == nil && matches.Length() > 0 && *matches.At(0) == "" {
		matches = matches.Slice(1, -1)
	}

	return array.Map(matches, func(v string, k int, input *array.Of[string]) string {
		return strings.ReplaceAll(v, "//", "/")
	}), nil
}
