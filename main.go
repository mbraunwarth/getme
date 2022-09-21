package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO Add recursion flag (greb -R ... / greb --no-recursion ...) to turn off explicit folder recursion
// TODO Add help/usage information
// TODO Time benchmark each function separately

func main() {
	// get file name(s) from command line argument
	args := os.Args[1:]
	// no args -> run greb in current working directory (i.e. `greb .`)
	if len(args) == 0 {
		args = []string{"."}
	}

	fileNames := getFileNames(args)

	for _, fname := range fileNames {
		// open file for reading
		f, err := os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// process file
		tags := readTagsFromFile(fname, f)

		// format output and print to stdout only if any tags are present
		if len(tags) != 0 {
			output := formatOutput(tags)
			fmt.Println(output)
		}
	}
}

// getFileNames reads the given arguments and returns a list of files names from which
// tags need to be grebbed.
func getFileNames(args []string) []string {
	fileNames := make([]string, 0)

	for _, arg := range args {
		// get fs.FileInfo object to check if the file is a directory
		info, err := os.Stat(arg)
		if err != nil {
			log.Fatal("X", err)
		}

		// carry over full path name relative from cwd
		parent := arg
		if info.IsDir() {
			// if current arg is a directory handle each subsequent file
			es, err := os.ReadDir(arg)
			if err != nil {
				log.Fatal(err)
			}

			args := make([]string, 0)
			for _, e := range es {
				// need to join relative file path because fs.DirEntry.Name() only
				// cares about the base name of a file
				args = append(args, filepath.Join(parent, e.Name()))
			}

			// opt for a depth first strategy
			fileNames = append(fileNames, getFileNames(args)...)

			continue
		}

		// if the current file is not a directory, just append to list of file names
		fileNames = append(fileNames, arg)
	}

	return fileNames
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
