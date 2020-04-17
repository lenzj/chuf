package main

import (
	"flag"
	"fmt"
	"io"
	"log"
        "git.lenzplace.org/lenzj/chunkio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Parse(in io.Reader, out io.Writer, start, end []byte, filtSlice []string) error {
	cio := chunkio.NewReader(in)
	for cio.GetErr() == nil {
		var cmd *exec.Cmd
		if len(filtSlice) > 1 {
			cmd = exec.Command(filtSlice[0], filtSlice[1:]...)
		} else {
			cmd = exec.Command(filtSlice[0])
		}
		cmd.Stdin = cio
		cmd.Stdout = out
		cio.SetKey(start)
		_, err := io.Copy(out, cio)
		if err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			return err
		}
		cio.Reset()
		cio.SetKey(end)
		if err = cmd.Run(); err != nil {
			return err
		}
		cio.Reset()
	}
	return nil
}

var Version string

func main() {
        // Get application name as executed from command prompt
        appName := filepath.Base(os.Args[0])

	// Set up formatting for error messages
	log.SetFlags(0)
	log.SetPrefix(appName + ": ")

	// Parse command line
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s BEGIN END FILTER\n"+
				"Transform sections of stdin by passing through a filter program.\n"+
				"Sections are defined by user selected BEGIN and END identifiers.\n"+
                                "Options:\n", appName)
		flag.PrintDefaults()
	}

	var vflag = flag.Bool("v", false, "display " + appName + " version")

	flag.Parse()

        // Display application version if requested
        if *vflag {
                fmt.Println(appName + " " + Version)
                os.Exit(0)
        }

	var (
		input              io.Reader
		output             io.Writer
		err                error
		filtStart, filtEnd string
		filtSlice          []string
	)

	switch flag.NArg() {
	case 3:
		filtStart = flag.Arg(0)
		if len(filtStart) == 0 {
			log.Fatalln("BEGIN chunk identifier must have length greater than zero")
		}
		filtEnd = flag.Arg(1)
		if len(filtEnd) == 0 {
			log.Fatalln("END chunk identifier must have length greater than zero")
		}
		filtSlice = strings.Split(flag.Arg(2), " ")
		if _, err = exec.LookPath(filtSlice[0]); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Println("incorrect number of arguments")
		flag.Usage()
		os.Exit(1)
	}

	input = os.Stdin
	output = os.Stdout

	err = Parse(input, output, []byte(filtStart), []byte(filtEnd), filtSlice)
	if err != nil {
		log.Fatalln(err)
	}
}
