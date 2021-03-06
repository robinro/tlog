package main

import (
	"bufio"
	"fmt"
	"github.com/jehiah/go-strftime"
	"github.com/ogier/pflag"
	"os"
	"time"
)

func main() {
	// Parsing command line options:
	var relative bool
	var incremental bool
	pflag.BoolVarP(&relative, "relative", "r", false, "Measure relative time")
	pflag.BoolVarP(&incremental, "incremental", "i", false, "Measure incremental time")
	pflag.Parse()
	if relative && incremental {
		fmt.Println("The options `relative` (`r`) and `incremental` (`i`) conflict.")
		fmt.Println("Exiting.")
		os.Exit(1)
	}

	// Defining variables that can be used in blocks
	var startTime time.Time
	var prevTime time.Time
	var format string

	if relative {
		startTime = time.Now()
	} else if incremental {
		prevTime = time.Now()
	} else {
		format = "%Y-%m-%dT%H:%M:%S%L"
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading stdin:", err)
		}
		if relative {
			fmt.Printf("%2.3f %s\n", time.Since(startTime).Seconds(), text)
		} else if incremental {
			fmt.Printf("%2.3f %s\n", time.Since(prevTime).Seconds(), text)
			prevTime = time.Now()
		} else {
			now := time.Now()
			fmt.Printf("%s %s\n", strftime.Format(format, now), text)
		}
	}
}
