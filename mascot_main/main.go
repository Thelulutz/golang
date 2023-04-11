package main

import (
	"fmt"

	"github.com/Thelulutz/golang-beginning/mascot_main/mascot"
	"rsc.io/quote"
)

func main() {
	// prints the mascot variable
	fmt.Println(mascot.BestMascot())
	// prints a quote from a package
	fmt.Println(quote.Go())

}
