package infr

import (
	"math"

	"github.com/takatoh/fft"
)

func Integrate(x []float64, n int) []float64 {
	nn := 2
	for nn < n {
		nn *= 2
	}

	var c []complex128
	for i := 0; i < n; i++ {
		c = append(c, complex(x[i], 0.0))
	}
	for i := n; i < nn; i++ {
		c = append(c, complex(0.0, 0.0))
	}

	cf := fft.FFT(c, nn)
	cf = integrate(cf, nn)
	c = fft.IFFT(cf, nn)

	var y []float64
	for i := 0; i < n; i++ {
		y = append(y, real(c[i]))
	}

	return y
}

func integrate(c []complex128, nn int) []complex128 {
	return c
}
