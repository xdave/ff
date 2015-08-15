// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package ff_test

import (
	"fmt"
	"testing"
)

import (
	. "github.com/smartystreets/goconvey/convey"
)

import (
	"github.com/xdave/ff"
)

var testKey = "foo"
var testValue = "bar"
var testKeys = []interface{}{testKey, 123, 5.0, true, false, nil}
var testValues = []interface{}{testValue, 123, 5.0, true, false, nil}

func TestNewParam(t *testing.T) {
	Convey("ff.NewParam(key, value interface{})", t, func() {
		for i, key := range testKeys {
			Convey(fmt.Sprintf("Given a '%T' key #%d", key, i), func() {
				param := ff.NewParam(key, nil)
				if fmt.Sprintf("%T", key) != "string" {
					Convey("Stored Key should be nil", func() {
						So(param.Key, ShouldBeNil)
					})
				} else {
					Convey("Stored Key should be a string", func() {
						So(param.Key, ShouldHaveSameTypeAs, "")
					})
				}
			})
		}
		for i, value := range testValues {
			Convey(fmt.Sprintf("With a '%T' value #%d", value, i), func() {
				param := ff.NewParam(testKey, value)
				Convey("Stored Value should same type", func() {
					So(param.Value, ShouldHaveSameTypeAs, value)
				})
			})
		}
	})
}

func TestParamSlice(t *testing.T) {
	Convey("ff.Param.Slice()", t, func() {
		for i, key := range testKeys {
			Convey(fmt.Sprintf("With a '%T' key #%d", key, i), func() {
				param := ff.NewParam(key, testValue)
				slice := param.Slice()
				if fmt.Sprintf("%T", key) != "string" {
					Convey("The len() slice should be less than 2", func() {
						So(len(slice), ShouldBeLessThan, 2)
					})
				} else {
					Convey("The len() of slice should be at least 1", func() {
						So(len(slice), ShouldBeGreaterThanOrEqualTo, 1)
					})
					Convey("The first element", func() {
						first := slice[0]
						Convey("Shouldn't be empty", func() {
							So(first, ShouldNotBeEmpty)
						})
						Convey("Should start with '-'", func() {
							So(first, ShouldStartWith, "-")
						})
						Convey("Should end with string of key", func() {
							So(first, ShouldEndWith, fmt.Sprintf("%v", key))
						})
					})
				}
			})
		}
		Convey("For both nil key and nil value", func() {
			param := ff.NewParam(nil, nil)
			slice := param.Slice()
			Convey("slice should be empty", func() {
				So(slice, ShouldBeEmpty)
				Convey("len() of slice should be 0", func() {
					So(len(slice), ShouldEqual, 0)
				})
				Convey("slice should resemble []string{}", func() {
					So(slice, ShouldResemble, []string{})
				})
			})
		})
		Convey("For nil key and non-nil value", func() {
			param := ff.NewParam(nil, testValue)
			slice := param.Slice()
			Convey("slice should not be empty", func() {
				So(slice, ShouldNotBeEmpty)
				Convey("len() of slice should be 1", func() {
					So(len(slice), ShouldEqual, 1)
				})
				Convey("slice should resemble []string{testValue}", func() {
					So(slice, ShouldResemble, []string{testValue})
				})
			})
		})
		Convey("For non-nil key and nil value", func() {
			param := ff.NewParam(testKey, nil)
			slice := param.Slice()
			Convey("slice should not be empty", func() {
				So(slice, ShouldNotBeEmpty)
				Convey("len() of slice should be 1", func() {
					So(len(slice), ShouldEqual, 1)
				})
				Convey("slice should resemble []string{-testKey}", func() {
					So(slice, ShouldResemble, []string{"-" + testKey})
				})
			})
		})
		Convey("For both non-nil key and value", func() {
			param := ff.NewParam(testKey, testValue)
			slice := param.Slice()
			Convey("slice should not be empty", func() {
				So(slice, ShouldNotBeEmpty)
				Convey("len() of slice should be 2", func() {
					So(len(slice), ShouldEqual, 2)
				})
				Convey("slice should resemble []string{-testKey, testValue}", func() {
					So(slice, ShouldResemble, []string{"-" + testKey, testValue})
				})
			})
		})
	})
}

func TestNewParamSet(t *testing.T) {
	Convey("ff.NewParamSet(...ff.Param)", t, func() {
		Convey("Result should NOT be nil", func() {
			set := ff.NewParamSet()
			So(set, ShouldNotBeNil)
		})
		Convey("ff.ParamSet.Len()", func() {
			Convey("If passed no params", func() {
				set := ff.NewParamSet()
				Convey("set.Len() should be 0", func() {
					So(set.Len(), ShouldEqual, 0)
				})
			})
			Convey("If passed a param with nil key and value", func() {
				param := ff.NewParam(nil, nil)
				set := ff.NewParamSet(param)
				Convey("set.Len() should be 0", func() {
					So(set.Len(), ShouldEqual, 0)
				})
			})
			Convey("If passed a param with nil key and non-nil value", func() {
				param := ff.NewParam(nil, testValue)
				set := ff.NewParamSet(param)
				Convey("set.Len() should be 1", func() {
					So(set.Len(), ShouldEqual, 1)
				})
			})
			Convey("If passed a param with non-nil key and nil value", func() {
				param := ff.NewParam(testKey, nil)
				set := ff.NewParamSet(param)
				Convey("set.Len() should be 1", func() {
					So(set.Len(), ShouldEqual, 1)
				})
			})
			Convey("If passed a param with non-nil key and non-nil value", func() {
				param := ff.NewParam(testKey, testValue)
				set := ff.NewParamSet(param)
				Convey("set.Len() should be 2", func() {
					So(set.Len(), ShouldEqual, 2)
				})
			})
			Convey("If passed 2 params with non-nil key and non-nil value", func() {
				param1 := ff.NewParam(testKey, testValue)
				param2 := ff.NewParam(testKey, testValue)
				set := ff.NewParamSet(param1, param2)
				Convey("set.Len() should be 4", func() {
					So(set.Len(), ShouldEqual, 4)
				})
			})
		})
		Convey("ff.ParamSet.Add(ff.Param)", func() {
			Convey("If passed a Param with len() 0", func() {
				param := ff.NewParam(nil, nil)
				set := ff.NewParamSet()
				set.Add(param)
				Convey("set.Len() should be 0", func() {
					So(set.Len(), ShouldEqual, 0)
				})
			})
			Convey("If passed a Param with len() 1", func() {
				param := ff.NewParam(testKey, nil)
				set := ff.NewParamSet()
				set.Add(param)
				Convey("set.Len() should be 1", func() {
					So(set.Len(), ShouldEqual, 1)
				})
			})
			Convey("If passed 2 Params with len() 1", func() {
				param1 := ff.NewParam(testKey, nil)
				param2 := ff.NewParam(nil, testValue)
				set := ff.NewParamSet()
				set.Add(param1, param2)
				Convey("set.Len() should be 2", func() {
					So(set.Len(), ShouldEqual, 2)
				})
			})
		})
		Convey("ff.ParamSet.Slice()", func() {
			Convey("Given 3 len(2) Params", func() {
				param1 := ff.NewParam(testKey, testValue)
				param2 := ff.NewParam(testKey, testValue)
				param3 := ff.NewParam(testKey, testValue)
				set := ff.NewParamSet(param1, param2)
				set.Add(param3)
				result := set.Slice()
				Convey("Calling Slice() should yield []string", func() {
					So(result, ShouldHaveSameTypeAs, []string{})
				})
				Convey("Result should not be empty", func() {
					So(result, ShouldNotBeEmpty)
				})
				Convey("set.Len() should equal number of actual items", func() {
					So(set.Len(), ShouldEqual, len(result))
				})
				Convey("result should resemble expected value", func() {
					So(result, ShouldResemble,
						[]string{
							"-" + testKey,
							testValue,
							"-" + testKey,
							testValue,
							"-" + testKey,
							testValue})
				})
			})
		})
	})
}

func TestBaseFile(t *testing.T) {
	Convey("Can create ff.BaseFile w/ Filename and Params", t, func() {
		filename := "test.mp4"
		param := ff.NewParam("foo", "bar")
		set := ff.NewParamSet(param)
		base := ff.BaseFile{Filename: filename, Params: set}
		Convey("It should contain the inital paramters", func() {
			So(base.Params.Slice(), ShouldContain, "-foo")
			So(base.Params.Slice(), ShouldContain, "bar")
		})
		Convey("ff.BaseFile.Name() returns the Filename", func() {
			So(base.Name(), ShouldEqual, filename)
		})
		Convey("ff.BaseFile.AddParam(ff.Param) adds the param", func() {
			newParam := ff.NewParam("baz", "bing")
			base.AddParam(newParam)
			So(base.Params.Slice(), ShouldContain, "-baz")
			So(base.Params.Slice(), ShouldContain, "bing")
		})
	})
}

func TestInput(t *testing.T) {
	filename := "test.mp4"
	param := ff.NewParam("foo", "bar")
	set := ff.NewParamSet(param)
	var input ff.File = ff.NewInput(filename, set)
	Convey("ff.NewInput(string, *ff.ParamSet)", t, func() {
		Convey("should not return nil", func() {
			So(input, ShouldNotBeNil)
		})
	})
	Convey("ff.File.Slice() should return paramters in correct order", t, func() {
		expected := []string{"-foo", "bar", "-i", filename}
		So(input.Slice(), ShouldResemble, expected)
	})
}

func TestOutput(t *testing.T) {
	filename := "test.mp4"
	param := ff.NewParam("foo", "bar")
	set := ff.NewParamSet(param)
	var output ff.File = ff.NewOutput(filename, set)
	Convey("ff.NewOutput(string, *ff.ParamSet)", t, func() {
		Convey("should not return nil", func() {
			So(output, ShouldNotBeNil)
		})
	})
	Convey("ff.File.Slice() should return paramters in correct order", t, func() {
		expected := []string{"-foo", "bar", filename, "-y"}
		So(output.Slice(), ShouldResemble, expected)
	})
}

func TestCommand(t *testing.T) {
	filename := "test.mp4"
	param := ff.NewParam("foo", "bar")
	set := ff.NewParamSet(param)
	var input ff.File = ff.NewInput(filename, set)
	var output ff.File = ff.NewOutput(filename, set)
	path := "ffmpeg"

	Convey("ff.NewCommand()", t, func() {
		Convey("Cannot call with empty command", func() {
			_, err := ff.NewCommand("", input, output)
			So(err, ShouldNotBeNil)
		})
		Convey("Can call with only input", func() {
			_, err := ff.NewCommand(path, input)
			So(err, ShouldBeNil)
		})
		Convey("Returns error with no input", func() {
			_, err := ff.NewCommand(path, nil, output)
			So(err, ShouldNotBeNil)
			_, err = ff.NewCommand(path, nil)
			So(err, ShouldNotBeNil)
		})
		Convey("Returns error no input or output", func() {
			_, err := ff.NewCommand(path, nil)
			So(err, ShouldNotBeNil)
		})
	})
	Convey("ff.Command.Slice() input+output", t, func() {
		cmd, _ := ff.NewCommand(path, input, output)
		slice := cmd.Slice()
		Convey("Should not return nil", func() {
			So(slice, ShouldNotBeNil)
		})
		Convey("Should return a []string slice", func() {
			So(slice, ShouldHaveSameTypeAs, []string{})
		})
		Convey("Should resemble expected value", func() {
			expected := []string{
				"-foo", "bar", "-i", filename,
				"-foo", "bar", filename, "-y",
			}
			So(slice, ShouldResemble, expected)
		})
	})
	Convey("ff.Command.Slice() input only", t, func() {
		cmd, _ := ff.NewCommand(path, input)
		slice := cmd.Slice()
		Convey("Should not return nil", func() {
			So(slice, ShouldNotBeNil)
		})
		Convey("Should return a []string slice", func() {
			So(slice, ShouldHaveSameTypeAs, []string{})
		})
		Convey("Should resemble expected value", func() {
			expected := []string{
				"-foo", "bar", "-i", filename,
			}
			So(slice, ShouldResemble, expected)
		})
	})
}

func TestInfo(t *testing.T) {
	var testJson = `{
        "streams": [{
            "codec_type": "video",
            "tags": { "rotate": "90" },
            "side_data_list": [{ "rotation": -90 }]
        },
        {
            "codec_type": "video",
            "tags": { "rotate": "180" },
            "side_data_list": [{ "rotation": -180 }]
        },
        { "codec_type": "audio" }]}`

	var info *ff.ProbeInfo
	var err error

	Convey("ff.NewInfo(string) should unmarshal json data", t, func() {
		info, err = ff.NewInfo(testJson)
		So(info, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("ff.Info.FilterStreams(ff.StreamType) (given test data)", t, func() {
		Convey("There should be 2 video streams", func() {
			vstreams := info.FilterStreams(ff.VideoStream)
			So(len(vstreams), ShouldEqual, 2)
		})
		Convey("There should be 1 audio stream", func() {
			astreams := info.FilterStreams(ff.AudioStream)
			So(len(astreams), ShouldEqual, 1)
		})
	})
	Convey("Video stream rotation", t, func() {
		vstream1 := info.FilterStreams(ff.VideoStream)[0]
		vstream2 := info.FilterStreams(ff.VideoStream)[1]
		Convey("ff.Stream.IsRotated() should tell if there's rotation", func() {
			rotated := vstream1.IsRotated()
			So(rotated, ShouldBeTrue)
		})
		Convey("First stream should be rotated by 90 degrees", func() {
			rotation := vstream1.Rotation()
			So(rotation, ShouldEqual, 90)
		})
		Convey("Second stream should be rotated by 180 degrees", func() {
			rotation := vstream2.Rotation()
			So(rotation, ShouldEqual, 180)
		})
	})
}
