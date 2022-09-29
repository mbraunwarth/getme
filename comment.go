package main

type Extension string

type Language struct {
	Name    string
	Comment string
}

type File struct {
	FullName string
	Base     string
	Path     string
	Ext      Extension
	Lang     Language
}

var langs = map[Extension]Language{
	Extension(".go"):   Language{"Golang", "//"},
	Extension(".rs"):   Language{"Rust", "//"},
	Extension(".java"): Language{"Java", "//"},
	Extension(".lua"):  Language{"Lua", "--"},
	Extension(".elm"):  Language{"Elm", "--"},
	Extension(".hs"):   Language{"Haskell", "--"},
	Extension(".py"):   Language{"Python", "#"},
	Extension(".rb"):   Language{"Ruby", "#"},
	Extension(".ex"):   Language{"Elixir", "#"},
	Extension(".exs"):  Language{"Elixir Script", "#"},
}
