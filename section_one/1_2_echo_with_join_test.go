package echo

import (
	"testing"
)

func BenchmarkEchoWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin()
	}
}
