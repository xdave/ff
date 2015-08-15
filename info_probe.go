//+build !test

// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package ff

import (
	"os/exec"
)

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
		return nil, err
	}

	out, err := exec.Command(cmdline.Path, cmdline.Slice()...).Output()
	if err != nil {
		return nil, err
	}

	info, err = NewInfo(string(out))
	if err != nil {
		return nil, err
	}

	return info, nil
}
