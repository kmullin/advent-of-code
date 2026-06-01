package cli

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/kmullin/advent-of-code/internal/aoc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const InputFilename = "input.txt"

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func readReadme() string {
	b, err := os.ReadFile("README.md")
	if err != nil {
		return ""
	}
	return string(b)
}

type ProblemRunner func([]byte) (any, error)

// NewCmd returns a new command for advent of code
func NewCmd(year, day int, funcs ...ProblemRunner) *cobra.Command {
	var flags struct {
		inputFilename string
	}
	var inputData []byte

	runPart := func(p int, f ProblemRunner) error {
		a, err := f(bytes.Clone(inputData))
		if err != nil {
			return fmt.Errorf("err running part %v: %w", p, err)
		}
		log.Info().Msgf("Part %v: %v", p, a)
		return nil
	}

	rootCmd := &cobra.Command{
		Use:           fmt.Sprintf("day%02v [part1 | part2]", day),
		Short:         fmt.Sprintf("Advent of Code %v day %02v", year, day),
		Long:          readReadme(),
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			config, err := initConfig(cmd)
			if err != nil {
				return fmt.Errorf("unable to init config: %w", err)
			}

			if config.Verbose {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}

			// attempt to read local file first, ignore non exist error
			inputData, err = os.ReadFile(flags.inputFilename)
			if err != nil && !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("err reading input file: %w", err)
			}

			// if we got nothing and we have a cookie
			if len(inputData) == 0 && config.SessionCookie != "" {
				aocclient := aoc.NewClient(config.SessionCookie, config.cacheDir)
				inputData, err = aocclient.GetInput(year, day)
				if err != nil {
					return err
				}
			}

			if len(inputData) == 0 {
				return fmt.Errorf("input data empty")
			}

			log.Info().Msgf("Advent of Code %v day %02v", year, day)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			for i, f := range funcs {
				if err := runPart(i+1, f); err != nil {
					return err
				}
			}
			return nil
		},
	}

	for i, f := range funcs {
		p := i + 1
		cmd := &cobra.Command{
			Use:     fmt.Sprintf("part%v", p),
			Short:   fmt.Sprintf("Run part%v", p),
			Aliases: strings.Split(fmt.Sprintf("p%v,%v", p, p), ","),
			RunE: func(cmd *cobra.Command, args []string) error {
				return runPart(p, f)
			},
		}
		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Debug logging")
	rootCmd.PersistentFlags().StringVarP(&flags.inputFilename, "input-file", "f", InputFilename, "File with input data")
	return rootCmd
}
