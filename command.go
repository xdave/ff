// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

// Package ff implements an ffmpeg (or ffprobe) command-line parameter builder
package ff

import (
	"fmt"
)

// Represents the entire commandline for ffmpeg/ffprobe
type Command struct {
	Path    string
	Input   File
	Outputs []File
}

// You can pass in only input (for ffprobe)
// You can pass in one or multiple ouptuts
// Returns a new Command structure
func NewCommand(path string, input File, outputs ...File) (cmd *Command, err error) {
	if path == "" {
		return nil, fmt.Errorf("Cannot create a Command with no path")
	}
	if input == nil {
		return nil, fmt.Errorf("Cannot create a Command with no input")
	}
	cmd = &Command{
		Path:  path,
		Input: input,
	}

	for _, output := range outputs {
		if output != nil {
			cmd.Outputs = append(cmd.Outputs, output)
		}
	}

	return cmd, nil
}

// Returns a []string slice of how the ffmpeg/ffprobe call should be represented
// Does not include the command Path before it
func (c *Command) Slice() (results []string) {
	results = []string{}
	results = append(results, c.Input.Slice()...)
	for _, output := range c.Outputs {
		results = append(results, output.Slice()...)
	}
	return
}
