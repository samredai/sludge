package models

type Row[T any] struct {
	Length int
	Data   T
}

type Collection[T any] struct {
	Items   []Row[T]
	Title   string
	Count   int
	Headers func(Row[T]) []string
	Row     func(Row[T]) []string
	RowData func(Collection[T]) []float64
}
