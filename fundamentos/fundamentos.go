package fundamentos

// Generics - Soma pode ser T, inteiro ou float
func Soma[T int | float64](a, b T) T {
	return a + b
}
