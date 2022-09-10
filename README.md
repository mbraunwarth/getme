# Greb

CLI utility to greb todos and fixmes in your code base.

## Usage

Use the `greb` utility on a file or a directory like so.

```sh
> cat main.go
package main

import "fmt"

func main() {
    // TODO write hello world
	fmt.Println("")
}
> greb main.go
./main.go:6:2: TODO -- write hello world
```

Without argument, `greb` scans the current directory recursively, directory recursion
can be suppressed with the `--no-recurse` or `-R` flag.

```sh
> greb
> ...
```

## Output

The output is parser friendly and can be used easily by editors. File name (relative
by default), line and column are separated by a colon followed by the [tag type](#tag-type)
and the tags message which are separated by a double dash.
