package file

import (
	"fmt"
	"os"

	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/wav"
)

type Read struct {
	path    string
	content []float32
}

// NewReadFile — create a new file struct
func NewReadFile(path string) (*Read, error) {
	file, openErr := os.Open(path)
	if openErr != nil {
		return nil, fmt.Errorf("open wav file %v: %w", path, openErr)
	}
	defer file.Close()

	dec := wav.NewDecoder(file)
	buf, decodeErr := dec.FullPCMBuffer()
	if decodeErr != nil {
		return nil, fmt.Errorf("decode wav file %v: %w", path, decodeErr)
	}

	// TODO: remove whisper.SampleRate
	if dec.SampleRate != whisper.SampleRate {
		return nil, fmt.Errorf("unsupported sample rate: %d", dec.SampleRate)
	}

	if dec.NumChans != 1 {
		return nil, fmt.Errorf("unsupported number of channels: %d", dec.NumChans)
	}

	return &Read{
		path:    path,
		content: buf.AsFloat32Buffer().Data,
	}, nil
}

// Path — get the path of the file
func (f *Read) Path() string {
	return f.path
}

// Content — get the content of the file
func (f *Read) Content() []float32 {
	return f.content
}
