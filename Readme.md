

# Brahms

Brahms is a command line utility written in Go that quickly generates good-looking visuals of midi-files.

It is inspired by [Off the staff](https://www.c82.net/offthestaff/) by Nicholas Rougeux.

Feel free to open PRs, issues, etc. I'm glad to help.

## Example

Below you can see an example of Beethoven's Symphony Nr. 9 (at least i think so, I can't remember exactly 🤷‍):

![Img](./midi2csv/beethoven_9.png)

Generated some other nice visualisation? Share it [here](https://github.com/Wulfheart/brahms/issues/new?assignees=&labels=showcase&template=showcase.md&title=) with the world in the project's showcase.

## Installation

### Prerequisites

Please ensure that you have installed [midicsv](https://www.fourmilab.ch/webtools/midicsv/) and made it available in your ``PATH`` variable so you can call it directly from the command-line.

You can test your midicsv installation by calling ``midicsv -u`` in your console. The expected output has to be similar to this:

```
Usage: Your\Path\Where\It\Is\Midicsv.exe [ options ] [ midi_file ] [ csv_file ]
       Options:
           -u            Print this message
           -v            Verbose: dump header and track information
Version 1.1 (January 2008)
```

### with go tools (works on all OS)

```
go install github.com/Wulfheart/brahms
```

### Windows

1. Download the latest ``.exe`` release [here](https://github.com/Wulfheart/brahms/releases).
2. Add the downloaded executable to your ``PATH`` variable.

## Usage

````
NAME:
   brahms - visualize your music

USAGE:
   brahms -i path/to/midi [global options]

VERSION:
   v0.1.0

GLOBAL OPTIONS:
   --in infile, -i infile     infile
   --out outfile, -o outfile  outfile
   --colors hex, -c hex       case-insensitive comma-separated string of hexcolors, e.g. '#40e0d0,#ff8c00,#ff0080'
   --fill-opacity value       (default: 0.5)
   --help, -h                 show help (default: false)
   --version, -v              print the version (default: false)
````

## Tips

For nice gradients look at https://uigradients.com/ or at https://webgradients.com/.


