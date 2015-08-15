package ff

import (
	"encoding/json"
	"strconv"
)

// Quick integer abs function (math package doesn't use ints)
func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

// Represents either Format or Stream tags embedded in the file
type Tags map[string]interface{}

// Container Format info
type Format struct {
	BitRate     string `json:"bit_rate"`
	Duration    string `json:"duration"`
	Filename    string `json:"filename"`
	LongName    string `json:"format_long_name"`
	Name        string `json:"format_name"`
	NumPrograms int    `json:"nb_programs"`
	NumStreams  int    `json:"nb_streams"`
	ProbeScore  int    `json:"probe_score"`
	Size        string `json:"size"`
	StartTime   string `json:"start_time"`
	Tags        Tags   `json:"tags"`
}

// Stream disposition
type Disposition struct {
	AttachedPic     int `json:"attached_pic"`
	CleanEffects    int `json:"clean_effects"`
	Comment         int `json:"comment"`
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	Karaoke         int `json:"karaoke"`
	Lyrics          int `json:"lyrics"`
	Original        int `json:"original"`
	VisualImpaired  int `json:"visual_impaired"`
}

// Stream side data. More fields may be added later as needed.
type SideData struct {
	Displaymatrix string `json:"displaymatrix"`
	Rotation      int    `json:"rotation"`
	Size          int    `json:"side_data_size"`
	Type          string `json:"side_data_type"`
}

// Used for enumerating the CodecType field in Stream
type StreamType string

const (
	VideoStream StreamType = "video"
	AudioStream            = "audio"
)

// Represents any kind of stream (Audio, Video, etc)
type Stream struct {
	AvgFrameRate       string      `json:"avg_frame_rate"`
	BitRate            string      `json:"bit_rate"`
	BitsPerRawSample   string      `json:"bits_per_raw_sample"`
	BitsPerSample      int         `json:"bits_per_sample"`
	ChannelLayout      string      `json:"channel_layout"`
	Channels           int         `json:"channels"`
	ChromaLocation     string      `json:"chroma_location"`
	CodecLongName      string      `json:"codec_long_name"`
	CodecName          string      `json:"codec_name"`
	CodecTag           string      `json:"codec_tag"`
	CodecTagString     string      `json:"codec_tag_string"`
	CodecTimeBase      string      `json:"codec_time_base"`
	CodecType          StreamType  `json:"codec_type"`
	CodedHeight        int         `json:"coded_height"`
	CodedWidth         int         `json:"coded_width"`
	ColorPrimaries     string      `json:"color_primaries"`
	ColorRange         string      `json:"color_range"`
	ColorSpace         string      `json:"color_space"`
	ColorTransfer      string      `json:"color_transfer"`
	DisplayAspectRatio string      `json:"display_aspect_ratio"`
	DivxPacked         string      `json:"divx_packed"`
	DmixMode           string      `json:"dmix_mode"`
	Duration           string      `json:"duration"`
	DurationTs         int         `json:"duration_ts"`
	HasBFrames         int         `json:"has_b_frames"`
	Height             int         `json:"height"`
	ID                 string      `json:"id"`
	Index              int         `json:"index"`
	IsAvc              string      `json:"is_avc"`
	Level              int         `json:"level"`
	LoroCmixlev        string      `json:"loro_cmixlev"`
	LoroSurmixlev      string      `json:"loro_surmixlev"`
	LtrtCmixlev        string      `json:"ltrt_cmixlev"`
	LtrtSurmixlev      string      `json:"ltrt_surmixlev"`
	MaxBitRate         string      `json:"max_bit_rate"`
	NalLengthSize      string      `json:"nal_length_size"`
	NumFrames          string      `json:"nb_frames"`
	PixFmt             string      `json:"pix_fmt"`
	Profile            string      `json:"profile"`
	QuarterSample      string      `json:"quarter_sample"`
	FrameRate          string      `json:"r_frame_rate"`
	Refs               int         `json:"refs"`
	SampleAspectRatio  string      `json:"sample_aspect_ratio"`
	SampleFmt          string      `json:"sample_fmt"`
	SampleRate         string      `json:"sample_rate"`
	StartPts           int         `json:"start_pts"`
	StartTime          string      `json:"start_time"`
	TimeBase           string      `json:"time_base"`
	Timecode           string      `json:"timecode"`
	Width              int         `json:"width"`
	Tags               Tags        `json:"tags"`
	Disposition        Disposition `json:"disposition"`
	SideDataList       []SideData  `json:"side_data_list"`
}

// Does what it says.
func (s Stream) IsRotated() bool {
	return s.Rotation() != 0
}

// Gets rotation value of video stream, either from stream side-data, or tags.
// Returns 0 if it's not rotated, or if we can't figure it out
func (s Stream) Rotation() int {
	for _, sdata := range s.SideDataList {
		return abs(sdata.Rotation)
	}
	rotationTag, ok := s.Tags["rotate"]
	if ok {
		rotation, ok := rotationTag.(string)
		if ok {
			val, _ := strconv.Atoi(rotation)
			return val
		}
	}
	return 0
}

// Just a nice structure that represents the JSON data returned by ffprobe.
// Mostly auto-generated with http://mholt.github.io/json-to-go
type ProbeInfo struct {
	Format  Format   `json:"format"`
	Streams []Stream `json:"streams"`
}

// Returns a new ProbeInfo structure from the input JSON
func NewInfo(jsonData string) (*ProbeInfo, error) {
	info := new(ProbeInfo)
	err := json.Unmarshal([]byte(jsonData), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// Filters out streams of the given StreamType
// Returns a []Stream slice, even if it's empty
func (info ProbeInfo) FilterStreams(t StreamType) (streams []Stream) {
	streams = []Stream{}
	for _, stream := range info.Streams {
		if stream.CodecType == t {
			streams = append(streams, stream)
		}
	}
	return
}
