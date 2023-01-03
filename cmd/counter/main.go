package main

import (
	"fmt"
	count "github.com/luisya22/the-power-of-go-tools-challenges"
)

func main() {
	c := count.NewCounter()

	fmt.Println(c.Next(), c.Next(), c.Next())

	cInf := count.NewCounter()

	cInf.Run()
}
