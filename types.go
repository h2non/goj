package goj

type T interface{}
type FilterFn func(T) bool
type Function func(T) T
