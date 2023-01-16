package main

import (
	"fmt"
	"os"
	count "the-power-of-go-tools-challenges"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(count.Lines())
}
