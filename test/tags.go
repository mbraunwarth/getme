package main

import (
	"log"
	"strings"
)

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

type Tag struct {
	FileName string
	TType    TagType
	Body     string
	Line     int
	Column   int
}

// TODO comment
func toTag(fname, line string, lineNumber int) Tag {
	// cut off line up to start of comment, leaving just the comment body
	ss := strings.SplitAfter(line, "//")
	// length of cut off part used as column parameter
	col := len(ss[0]) + 1
	comment := strings.TrimSpace(ss[1])

	// separate the tag name from the tags body: `TAGNAME body` -> `TAGNAME`, `body`
	// BUG hitting a line with a `//` followed by nothing or not enough words,
	//	   i.e. at least 2 words separated by a white space, the program panics
	//     with `index out of bounds` as the splitted slice is expected to have
	//     at least two entries
	tagPartsRaw := strings.SplitN(comment, " ", 2)
	log.Printf("%+v\n", tagPartsRaw)
	// parse tag type from tag name
	ttype := typeFromName(tagPartsRaw[0])
	body := tagPartsRaw[1]

	return Tag{fname, ttype, body, lineNumber, col}
}
