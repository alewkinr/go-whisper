package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alewkinr/go-whisper/cmd"
	"github.com/urfave/cli/v2"
)

const (
	appName        = "go-whisper"
	appDescription = "A simple CLI tool for OpenAI whisper usage"
)

func main() {
	commands := make([]*cli.Command, 0)
	commands = append(commands, cmd.Transcribe)

	app := &cli.App{
		Name:  appName,
		Usage: appDescription,
		Action: func(*cli.Context) error {
			fmt.Println("Choose any command. Write --help for more info.")
			return nil
		},
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
