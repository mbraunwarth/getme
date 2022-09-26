# Getme

CLI utility to find todos, fixmes etc. in your code base. Those tags will referred to
as *comment tags*.

## Usage

Use the `getme` utility on a file or a directory like so.

```sh
> cat main.go
package main

import "fmt"

func main() {
    // TODO write hello world
	fmt.Println("")
}
> getme main.go
./main.go:6:2: TODO -- write hello world
```

Without argument, `getme` scans the current directory recursively.

## Output

The output is parser friendly and can be used easily by editors. File name (relative
by default), line and column are separated by a colon followed by the 
and the tags message which are separated by a double dash.
