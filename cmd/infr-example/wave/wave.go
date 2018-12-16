package wave

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"math"
)

type Wave struct {
	Name string
	Dt   float64
	Data []float64
}

func newWave() *Wave {
	p := new(Wave)
	return p
}

func LoadCSV(filename string) *Wave {
	var reader *csv.Reader
	var row []string
	var err error
	var wave *Wave
	var t1, t2, d float64
	var data []float64

	wave = newWave()
	t1 = 0.0
	t2 = 0.0

	read_file, _ := os.OpenFile(filename, os.O_RDONLY, 0600)
	reader = csv.NewReader(read_file)

	row, err = reader.Read()
	wave.Name = row[1]
	for {
		row, err = reader.Read()
		if err == io.EOF {
			wave.Dt = round(t2 - t1, 2)
			wave.Data = data
			return wave
		}
		t1 = t2
		t2, _ = strconv.ParseFloat(row[0], 64)
		d, _ = strconv.ParseFloat(row[1], 64)
		data = append(data, d)
	}
}

func round(val float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}