package main

import (
	"fmt"
	"os"

	"github.com/takatoh/fft"
	"github.com/takatoh/infr"
	"github.com/takatoh/infr/cmd/infr-example/wave"
)

func main() {
	csvfile := os.Args[1]
	w := wave.LoadCSV(csvfile)
	acc := w.Data
	n := len(acc)

	nn := 2
	for nn < n {
		nn *= 2
	}

	var c []complex128
	for i := 0; i < n; i++ {
		c = append(c, complex(acc[i], 0.0))
	}
	for i := n; i < nn; i++ {
		c = append(c, complex(0.0, 0.0))
	}
	cf := fft.FFT(c, nn)
	cf = infr.Integrate(cf, nn)
	c = fft.IFFT(cf, nn)

	for i := 0; i < n; i++ {
		fmt.Printf("%.3f\n", real(c[i]))
	}
}
