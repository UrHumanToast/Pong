package utilities

type rangeTypes interface {
	float64 | int64 | uint64
}

func InRange[T rangeTypes](min T, val T, max T) bool {
	return (min < val) && (val < max)
}

func OnRange[T rangeTypes](min T, val T, max T) bool {
	return (min <= val) && (val <= max)
}
