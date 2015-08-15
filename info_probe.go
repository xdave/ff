// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package ff

import (
	"os/exec"
)

// Used for faking exec.Cmd.Output() to test Probe()
type CommandInterface interface {
	Output() ([]byte, error)
}

// Used for faking exec.Command() to test Probe()
type CommandFunc func(string, ...string) CommandInterface

// Used for faking exec.Command() to test Probe()
var DefaultCommandFunc CommandFunc = func(name string, arg ...string) CommandInterface {
	return exec.Command(name, arg...)
}

// Actually calls ffprobe on a file and returns a ProbeInfo object
func Probe(path string) (info *ProbeInfo, err error) {
	input := NewInput(
		path,
		NewParamSet(
			NewParam("v", "quiet"),
			NewParam("print_format", "json"),
			NewParam("show_format", nil),
			NewParam("show_streams", nil),
		),
	)

	cmdline, err := NewCommand("ffprobe", input)
	if err != nil {
		return
	}

	var cmd CommandInterface
	cmd = DefaultCommandFunc(cmdline.Path, cmdline.Slice()...)
	out, err := cmd.Output()
	if err != nil {
		return
	}

	return NewInfo(string(out))
}
