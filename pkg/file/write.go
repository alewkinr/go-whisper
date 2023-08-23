package file

import (
	"fmt"
	"os"
)

// Write — filer to write result
type Write struct {
	path string
	file *os.File
}

// NewWriteFile — create a new file struct to write result
func NewWriteFile(path string) (*Write, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create file %v: %w", path, err)
	}

	return &Write{
		path: path,
		file: f,
	}, nil
}

// Write — implement io.Writer
func (w Write) Write(p []byte) (int, error) {
	return w.writeString(string(p))
}

// WriteString — write string to file
func (w Write) writeString(text string) (int, error) {
	l, err := w.file.WriteString(text)
	if err != nil {
		return 0, fmt.Errorf("write string to file %v: %w", w.path, err)
	}
	return l, nil
}

// Close — close file
func (w Write) Close() {
	if w.file != nil {
		_ = w.file.Close()
	}
}
