package main

import (
	"fmt"
	"os"

	"github.com/takatoh/infr"
	"github.com/takatoh/infr/cmd/infr-example/wave"
)

func main() {
	csvfile := os.Args[1]
	w := wave.LoadCSV(csvfile)
	acc := w.Data

	vel := infr.Integrate(acc, len(acc))
	n := len(vel)
	for i := 0; i < n; i++ {
		fmt.Printf("%.3f\n", vel[i])
	}
}
