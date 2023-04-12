package mascot_test

import (
	"testing"

	"github.com/Thelulutz/golang-beginning/mascot_main/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
