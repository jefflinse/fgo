// Package fgo provides a set of function-style slice manipulation functions.
package fgo

// Filter returns a new slice containing only elements that satisfy the predicate f.
func Filter[A any](items []A, f func(A) bool) []A {
	out := []A{}
	for _, x := range items {
		if f(x) {
			out = append(out, x)
		}
	}

	return out
}

// Map returns a new slice with the results of applying f to each element of s.
func Map[A any, B any](items []A, f func(A) B) []B {
	out := make([]B, len(items))
	for i, x := range items {
		out[i] = f(x)
	}

	return out
}

// Reduce applies f to each element of s, starting with an initial value, and returns the result.
func Reduce[ElemType any, ValueType any](items []ElemType, f func(ValueType, ElemType) ValueType, initial ValueType) ValueType {
	out := initial
	for _, x := range items {
		out = f(out, x)
	}

	return out
}
