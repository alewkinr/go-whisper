package cmd

import (
	"fmt"

	"github.com/alewkinr/go-whisper/pkg/whisper"
	"github.com/urfave/cli/v2"
)

const (
	transcribeCmdName = "transcribe"
)

// Transcribe — represents the transcribe command
var Transcribe = &cli.Command{
	Name: transcribeCmdName,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "Load model for text transcription from `FILE`",
		},
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "Load audio for text transcription from `FILE`",
		},
		&cli.StringFlag{
			Name:    "language",
			Value:   "RU",
			Aliases: []string{"l"},
			Usage:   "Language used in audio file",
		},
		&cli.StringFlag{
			Name:    "out",
			Aliases: []string{"o"},
			Usage:   "Output file for text transcription",
		},
	},
	Action: transcribeAction,
}

// transcribeAction — represents the transcribe action
var transcribeAction = func(c *cli.Context) error {
	modelPath := c.String("model")
	filePath := c.String("file")
	lang := c.String("language")
	out := c.String("out")

	if modelPath == "" {
		return fmt.Errorf("%w: model path is required", ErrMissingFlag)
	}

	if filePath == "" {
		return fmt.Errorf("%w: model path is required", ErrMissingFlag)
	}

	model, loadModelErr := whisper.New(modelPath, lang, out)
	if loadModelErr != nil {
		return fmt.Errorf("load model: %w", loadModelErr)
	}
	defer model.Close()

	if processFile := model.Process(filePath); processFile != nil {
		return fmt.Errorf("process file: %w", processFile)
	}
	return nil
}
