package main

import (
	"testing"
)

func BenchmarkSimpleEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkEchoWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(uint64(i))
	}
}
