package main

import (
	"fmt"
	"os"

	tern "tern/internal"
)

func main() {
	res, err := tern.Run(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
	_, _ = fmt.Fprintln(os.Stdout, res)
}
