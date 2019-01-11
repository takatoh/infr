package infr

import (
	"math"
	"math/cmplx"
)

func Integrate(c []complex128, nn int, dt, v0 float64) []complex128 {
	s := make([]complex128, nn)
	nfold := nn / 2
	pn := math.Pi / float64(nn)

	cr := 2.0 * pn * v0 / dt
	for k := 1; k < nfold; k++ {
		cr = cr - 2.0 * imag(c[k]) / float64(k)
		s[k] = complex(pn, 0.0) * c[0] * complex(-1.0, 1.0 / math.Tan(float64(k) * pn)) -
			complex(0.0, 1.0 / float64(k)) * c[k]
		s[nn - k] = cmplx.Conj(s[k])
	}
	s[0] = complex(cr, 0.0) + complex(pn * float64(nn - 1), 0.0) * c[0]
	s[nfold] = complex(-1.0 * pn, 0.0) * c[0]

	for k := 0; k < nn; k++ {
		s[k] = complex(float64(nn) * dt / (2.0 * math.Pi), 0.0) * s[k] 
	}

	return s
}
