package mascot_test

import (
	"testing"

	"github.com/Thelulutz/golang-beginning/mascot"
)

func TestMascot(t *testing.T) {
    if mascot.BestMascot() != "Go Gopher" {
        t.Fatal("Wrong mascot :(")
    }
}