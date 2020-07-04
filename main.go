package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Wulfheart/brahms/score"
	"github.com/Wulfheart/brahms/viz"
	"github.com/mingrammer/cfmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			cfmt.Errorln(err)
			fmt.Println("")
		}
	}()

	app := &cli.App{
		Name:            "brahms",
		Usage:           "visualize your music",
		Version:         "v0.1.1",
		HideHelpCommand: true,
		UsageText:       "brahms -i path/to/midi [global options]",
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			cfmt.Errorln(err)
			cfmt.Infoln("Here again the help text. I hope it helps.🧙‍\n")
			cli.ShowAppHelp(context)
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "in",
				Aliases:  []string{"i"},
				Usage:    "`infile`",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "out",
				Aliases: []string{"o"},
				Usage:   "`outfile` ",
			},
			&cli.StringFlag{
				Name:    "colors",
				Aliases: []string{"c"},
				Usage:   "case-insensitive comma-separated string of `hex`colors, e.g. '#40e0d0,#ff8c00,#ff0080'",
			},
			&cli.Float64Flag{
				Name:   "width",
				Value:  800,
				Hidden: true,
			},
			&cli.Float64Flag{
				Name:   "height",
				Value:  800,
				Hidden: true,
			},
			&cli.Float64Flag{
				Name:   "max-r",
				Value:  350,
				Hidden: true,
			},
			&cli.Float64Flag{
				Name:   "min-r",
				Value:  0,
				Hidden: true,
			},
			&cli.Float64Flag{
				Name:   "max-r-node",
				Value:  15,
				Hidden: true,
			},
			&cli.Float64Flag{
				Name:  "fill-opacity",
				Value: 0.5,
				// Hidden: true,
			},
			&cli.StringFlag{
				Name:  "midi2csv",
				Value: "midicsv",
				Usage: "Provide a custom location for your midicsv command",
				// Hidden: true,
			},
		},
		Action: func(c *cli.Context) error {

			input, err := existingFilepath(c.String("in"), false)
			if err != nil {
				return err
			}
			sc := score.Read(input, score.ReadMidi, c.String("midi2csv"))
			buf := new(bytes.Buffer)
			// TODO: Make error handling for wrong formatted strings
			if strings.Contains(c.String("colors"), " ") {
				return fmt.Errorf("don't use a space in your hexcolor string")
			}
			colors, err := viz.MakePalette(len(sc), strings.Split(c.String("colors"), ",")...)
			colorsEmpty := c.String("colors") == ""
			viz.RenderCircle(buf, sc, viz.CircleConfig{
				MaxR:        c.Float64("max-r"),
				MinR:        c.Float64("min-r"),
				MaxRNode:    c.Float64("max-r-node"),
				Width:       c.Float64("width"),
				Height:      c.Float64("height"),
				FillOpacity: c.Float64("fill-opacity"),
				Colors:      colors,
				Filled:      !colorsEmpty,
				Stroke:      colorsEmpty,
			})

			// No output defined
			if c.String("out") == "" {
				f := bufio.NewWriter(os.Stdout)
				defer f.Flush()
				_, err = f.Write(buf.Bytes())
				if err != nil {
					return err
				}
			} else {
				out, err := existingFilepath(c.String("out"), true)
				if err != nil {
					return err
				}
				f, err := os.Create(out)
				if err != nil {
					return err
				}
				_, err = f.Write(buf.Bytes())
				if err != nil {
					return err
				}
				cfmt.Successf("Midi rendered as svg at %s\n", out)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}

func existingFilepath(p string, createIfNotExists bool) (path string, err error) {
	path, err = filepath.Abs(p)
	if err != nil {
		return "", err
	}
	// File does not exist
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) && createIfNotExists{
			_, e := os.Create(path)
			if e != nil {
				return "", e
			}
		} else {

		return "", fmt.Errorf("it seems like there was an issue with the file %s:\n%s", path, err)
		}
	}
	return path, nil
}
