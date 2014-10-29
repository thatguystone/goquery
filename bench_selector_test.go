package goquery

import (
	"testing"

	"code.google.com/p/cascadia"
)

func BenchmarkSelectorSimpleCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cascadia.MustCompile("#main")
	}
}

func BenchmarkSelectorComplexCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cascadia.MustCompile(".main .article ul > lu:first-child a.add")
	}
}

func BenchmarkSelectorCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getSelector(".main .article ul > lu:first-child a.add")
	}
}
