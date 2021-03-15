package main

import "testing"

func BenchmarkCountLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountLine("./test.txt")
	}
}

func BenchmarkCountWithScanner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountWithScanner("./test.txt")
	}
}
