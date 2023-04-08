package main

import (
	"fmt"

	"github.com/Thelulutz/golang-beginning/mascot"
	"github.com/Thelulutz/golang-beginning/variables"
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())

	variables.PrintVariables()
}
