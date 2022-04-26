package smi

type Rectangle struct {
	length, width float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.length + rec.width)
}

func Area(l, w float64) float64 {
	return l * w
}
