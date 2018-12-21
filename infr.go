package infr

import (
	"math"
	"math/cmplx"
)

func Integrate(c []complex128, nn int) []complex128 {
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
