package infr

import (
	"math"
	"math/cmplx"

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
	s := make([]complex128, nn)
	nfold := nn / 2
	pn := math.Pi / float64(nn)

	c0 := real(c[0])
	s0 := float64(nn) / 2.0 * c0 * pn
	for k := 1; k < nfold - 1; k++ {
		s0 = s0 - imag(c[k]) / float64(k)
		s[k] = complex(-1.0 * c0 * pn, c0 * pn / math.Tan(float64(k) * pn)) - complex(0.0, 1.0 / float64(k)) * c[k]
		s[nn - k] = cmplx.Conj(s[k])
	}
	s[0] = complex(s0 * 2.0, 0.0)
	s[nfold] = complex(-1.0 * c0 * pn, 0.0)

	return s
}
