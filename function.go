package goj

func combine(a Function, b Function) Function {
	return func(arg T) T {
		return a(b(arg))
	}
}

func Compose(first Function, f ...Function) Function {
	for _, v := range f {
		first = combine(first, v)
	}
	return first
}
