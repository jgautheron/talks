// mypackage_test.go
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40)
	}
}

func BenchmarkFibonacciFast(b *testing.B) {
	resetLookupTable()

	for i := 0; i < b.N; i++ {
		FibonacciFast(40)
	}
}