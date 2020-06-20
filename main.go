package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Name:            "brahms",
		Usage:           "visualize your music",
		Version:         "v0.1.0",
		HideHelpCommand: true,
		// ArgsUsage: " ",
		UsageText: "brahms -i path/to/midi [global options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "`infile`",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "`outfile` ",
			},
			&cli.StringFlag{
				Name:    "colors",
				Aliases: []string{"c"},
				Usage:   "language for the greeting",
			},
			&cli.IntFlag{
				Name:   "width",
				Value:  800,
				Hidden: true,
			},
			&cli.IntFlag{
				Name:   "height",
				Value:  800,
				Hidden: true,
			},
			&cli.IntFlag{
				Name:   "max-r",
				Value:  350,
				Hidden: true,
			},
			&cli.IntFlag{
				Name:   "min-r",
				Value:  0,
				Hidden: true,
			},
			&cli.IntFlag{
				Name:   "max-r-node",
				Value:  15,
				Hidden: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.Int("width"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Compiling finished")
	// // p, err := filepath.Abs("./midi2csv/Beethoven_9.mid")
	// // p, err := filepath.Abs("./midi2csv/bcs.mid")
	// // p, err := filepath.Abs("./midi2csv/bach_brandenburg_3.mid")
	// p, err := filepath.Abs("./midi2csv/messiah.mid")
	// if err != nil {
	// 	panic(err)
	// }
	// sc := score.Read(p, score.ReadMidi)
	// fmt.Println(sc.AvgDuration(), sc.MaxDuration(), sc.MinDuration())
	// buf := new(bytes.Buffer)
	// viz.CreateCircular(sc, buf)
	// p, err = filepath.Abs("./midi2csv/out.svg")
	// if err != nil {
	// 	panic(err)
	// }
	// f, err := os.Create(p)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = f.Write(buf.Bytes())
	// if err != nil {
	// 	panic(err)
	// }

}
