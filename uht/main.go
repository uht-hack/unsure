package main

import (
	"github.com/corverroos/unsure"
	"github.com/luno/jettison/errors"

	"github.com/uht-hack/unsure/state"
)

func main() {
	unsure.Bootstrap()

	_, err := state.New()
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	//loser_ops.StartLoops(s)

	unsure.WaitForShutdown()
}
