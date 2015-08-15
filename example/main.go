// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package main

import (
	"flag"
	"fmt"
	"os/exec"
)

import (
	"github.com/xdave/ff"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Usage: example INFILE")
		return
	}

	filename := flag.Args()[0]

	input := ff.NewInput(
		filename,
		ff.NewParamSet(
			ff.NewParam("v", "quiet"),
			ff.NewParam("print_format", "json"),
			ff.NewParam("show_format", nil),
			ff.NewParam("show_streams", nil),
		),
	)

	cmdline, err := ff.NewCommand("ffprobe", input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Calling", cmdline.Path, cmdline.Slice())

	out, err := exec.Command(cmdline.Path, cmdline.Slice()...).CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
