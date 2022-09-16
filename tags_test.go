package main

import (
	"testing"
)

// ignoring 'Undefined' tag
var stringTagMap = map[string]TagType{
	"TODO":  Todo,
	"FIXME": Fixme,
	"BUG":   Bug,
	"XXX":   Xxx,
}

func TestTagSupportCoverage(t *testing.T) {
	want := len(stringTagMap)
	got := int(TagsEnumSize) - 1

	// sub 1 for ignoring 'Undefined' here and 'TagsEnumSize' itself (iota starts counting at 0)
	if want != got {
		t.Fatalf("TagType enum changed => tags: %d, covered in tests: %d", want, got)
	}
}
func TestTypeFromName(t *testing.T) {
	want := stringTagMap

	names := []string{"TODO", "FIXME", "BUG", "XXX"}

	for _, name := range names {
		got := typeFromName(name)
		if got != want[name] {
			t.Fatalf(`typeFromName(%s) => want %s, got %s`, name, want[name], got)
		}
	}
}

func TestToTag(t *testing.T) {
	line := "// TODO add 'World' to return string"

	want := Tag{
		FileName: "script",
		TType:    Todo,
		Body:     "add 'World' to return string",
		Line:     2,
		Column:   3,
	}

	got := toTag("script", line, 2)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}
