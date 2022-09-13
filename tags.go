package main

import (
	"strings"
)

// The TagType type alias represents an enumeration of a comment tags type.
type TagType int

const (
	Todo TagType = iota
	Fixme
	Bug
	Xxx
	Undefined
)

func (tt TagType) String() string {
	switch tt {
	case Todo:
		return "TODO"
	case Fixme:
		return "FIXME"
	case Bug:
		return "BUG"
	case Xxx:
		return "XXX"
	default:
		return "Undefined"
	}
}

// typeFromName returns the TagType given its upper case string representation.
// E.g. the given the string `TODO`, the function returns the `Todo` TagType.
func typeFromName(name string) TagType {
	switch name {
	case "TODO":
		return Todo
	case "FIXME":
		return Fixme
	case "BUG":
		return Bug
	case "XXX":
		return Xxx
	default:
		return Undefined
	}
}

// The Tag struct representing a tag.
type Tag struct {
	FileName string
	TType    TagType
	Body     string
	Line     int
	Column   int
}

// toTag separates the given line in a tag type and its body and returns a Tag struct
// based on that and the given file name and line number.
func toTag(fname, line string, lineNumber int) Tag {
	// cut off line up to start of comment, leaving just the comment body
	ss := strings.SplitAfter(line, "//")
	// length of cut off part used as column parameter
	col := len(ss[0]) + 1
	comment := strings.TrimSpace(ss[1])

	// BUG #1 hitting a line with a `//` followed by nothing or not enough words,
	//	   i.e. at least 2 words separated by a white space, the program panics
	//     with `index out of bounds` as the splitted slice is expected to have
	//     at least two entries

	// separate the tag name from the tags body: `TAGNAME body` -> `TAGNAME`, `body`
	tagPartsRaw := strings.SplitN(comment, " ", 2)
	// parse tag type from tag name
	ttype := typeFromName(tagPartsRaw[0])
	body := tagPartsRaw[1]

	return Tag{fname, ttype, body, lineNumber, col}
}
