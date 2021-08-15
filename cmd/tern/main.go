package main

import (
	"fmt"
	"log"
	"os"

	tern "tern/internal"
)

func main() {
	res, err := tern.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
