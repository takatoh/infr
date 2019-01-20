package main

import (
	"fmt"
	"os"

	"github.com/takatoh/fft"
	"github.com/takatoh/infr"
	"github.com/takatoh/seismicwave"
)

func main() {
	csvfile := os.Args[1]
	w, e := seismicwave.LoadCSV(csvfile)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	acc := w[0].Data
	dt := w[0].Dt
	n := len(acc)
	v0 := acc[0] * dt

	nn := 2
	for nn < n {
		nn *= 2
	}

	c := make([]complex128, nn)
	for i := 0; i < n; i++ {
		c[i] = complex(acc[i], 0.0)
	}
	for i := n; i < nn; i++ {
		c[i] = complex(0.0, 0.0)
	}

	cf := fft.FFT(c, nn)
	cf = infr.Integrate(cf, nn, dt, v0)
	c = fft.IFFT(cf, nn)

	vel := make([]float64, n)
	for i := 0; i < n; i++ {
		vel[i] = real(c[i])
	}

	cf = fft.FFT(c, nn)
	cf = infr.Integrate(cf, nn, dt, 0.0)
	c = fft.IFFT(cf, nn)

	dis := make([]float64, n)
	for i := 0; i < n; i++ {
		dis[i] = real(c[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%f,%f\n", vel[i], dis[i])
	}
}
