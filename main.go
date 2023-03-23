package main

import (
	"fmt"
	"path/filepath"
)

// issue: https://github.com/golang/go/issues/59188

func main() {
	match, err := filepath.Match(`/etc/*`, "/etc/foo/bar")
	//match, err := filepath.Match(`\etc\*`, `\etc\foo\bar`)
	//match, err := filepath.Match(filepath.ToSlash(filepath.Clean(`\etc\*`)), filepath.ToSlash(filepath.Clean(`\etc\foo\bar`)))
	fmt.Println(match, err) // This should return the same output on any OS
}
