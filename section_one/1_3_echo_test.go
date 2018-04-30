package echo

import (
	"testing"
)

func BenchmarkSimpleEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simpleEcho()
	}
}

func BenchmarkEchoWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin()
	}
}
