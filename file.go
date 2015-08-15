// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package ff

// Interface for either an input file or an output file
type File interface {
	Name() string
	AddParam(...Param)
	Slice() []string
}

// Base structure representing an input file or an output file.
// Not meant to be used directly; use InputFile or OutputFile
type BaseFile struct {
	Filename string
	Params   *ParamSet
}

// Add a parameter to this file
func (f BaseFile) AddParam(param ...Param) {
	f.Params.Add(param...)
}

// Returns the filename passed into the constructor
func (f BaseFile) Name() string {
	return f.Filename
}

// Represents an input file for ffmpeg/ffprobe
type InputFile struct {
	BaseFile
}

// Returns a slice representation of how an input file should be passed to
// ffmpeg/ffprobe
func (f InputFile) Slice() (result []string) {
	result = append(result, f.Params.Slice()...)
	result = append(result, "-i", f.Filename)
	return
}

// Returns a new File for input to ffmpeg/ffprobe
func NewInput(filename string, params *ParamSet) File {
	return &InputFile{
		BaseFile: BaseFile{
			Filename: filename,
			Params:   params,
		},
	}
}

// Represents an output file for ffmpeg/ffprobe
type OutputFile struct {
	BaseFile
}

// Returns a slice representation of how an output file should be passed to
// ffmpeg/ffprobe
func (f OutputFile) Slice() (result []string) {
	result = append(result, f.Params.Slice()...)
	result = append(result, f.Filename, "-y")
	return
}

// Returns a new File for output from ffmpeg/ffprobe
func NewOutput(filename string, params *ParamSet) File {
	return &OutputFile{
		BaseFile: BaseFile{
			Filename: filename,
			Params:   params,
		},
	}
}
