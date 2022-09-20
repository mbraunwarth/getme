package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// get file name from command line argument
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("exactly one file name required")
	}

	// open file for reading
	fname := args[0]
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	// process file
	tags := readTagsFromFile(fname, f)

	// format output and print to stdout
	output := formatOutput(tags)

	fmt.Println(output)
}

// readTagsFromFile scans through the given reader line by line and filters out
// the lines with comment tags. The function returns a slice of tags contained in
// the scanned file.
func readTagsFromFile(fname string, rd io.Reader) []Tag {
	tags := make([]Tag, 0)
	lineNumber := 0
	sc := bufio.NewScanner(rd)
	for sc.Scan() {
		// skim l for comments
		line := sc.Text()
		lineNumber = lineNumber + 1

		p := compiledTagRegexp()
		// match the line against generated regexp
		if p.MatchString(line) {
			// create tag from comment and append to tag list
			t := toTag(fname, line, lineNumber)
			tags = append(tags, t)
		}
	}

	return tags
}

// formatOutput produces a editor friendly output format from a list of tags.
// Example output: `./main.go:6:2: TAGNAME -- body string`
func formatOutput(tags []Tag) string {
	var sb strings.Builder

	for i, t := range tags {
		line := fmt.Sprintf(
			"%s:%d:%d: %s -- %s",
			t.FileName,
			t.Line,
			t.Column,
			t.TType,
			t.Body,
		)
		sb.WriteString(line)

		if i < len(tags)-1 {
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}
