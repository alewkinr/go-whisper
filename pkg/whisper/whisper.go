package whisper

import (
	"fmt"
	"io"
	"os"

	"github.com/alewkinr/go-whisper/pkg/file"
	timetool "github.com/alewkinr/go-whisper/pkg/time"
	cppWhisper "github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
)

// Whisper — wrapper for cppWhisper.Model
type Whisper struct {
	modelPath string
	lang      string
	out       io.Writer

	model cppWhisper.Model
}

// New — creates new Whisper
func New(modelPath, lang, out string) (*Whisper, error) {
	var writer io.Writer
	var createOutErr error

	switch {
	case out != "":
		writer, createOutErr = file.NewWriteFile(out)
	default:
		writer = os.Stdout
	}

	if createOutErr != nil {
		return nil, fmt.Errorf("create out file: %w", createOutErr)
	}

	model, loadModelErr := cppWhisper.New(modelPath)
	if loadModelErr != nil {
		return nil, fmt.Errorf("load model: %w", loadModelErr)
	}

	return &Whisper{
		modelPath: modelPath,
		lang:      lang,
		out:       writer,
		model:     model,
	}, nil
}

// Close — closes model
func (w Whisper) Close() error {
	return w.model.Close()
}

// Process — processes file
func (w Whisper) Process(filePath string) error {
	ctx, err := w.model.NewContext()
	if err != nil {
		return fmt.Errorf("create context: %w", err)
	}

	f, createFile := file.NewReadFile(filePath)
	if createFile != nil {
		return fmt.Errorf("create file: %w", createFile)
	}

	if processErr := ctx.Process(f.Content(), w.segmentHandler, nil); processErr != nil {
		return fmt.Errorf("process file: %w", processErr)
	}

	return nil
}

// segmentHandler — handler for segment
func (w Whisper) segmentHandler(s cppWhisper.Segment) {
	text := fmt.Sprintf("[TIME] %s --> %s\n[TEXT] %s\n", timetool.TimestampToSrt(s.Start), timetool.TimestampToSrt(s.End), s.Text)

	_ = w.writeTo(w.out, text)
}

// writeTo — writes text to out
func (w Whisper) writeTo(out io.Writer, text string) error {
	_, err := fmt.Fprintln(out, text)
	return err
}
