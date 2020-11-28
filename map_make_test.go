package goplayground

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m map[string]int

func BenchmarkMap(b *testing.B) {
	for length, builder := range buildersRegistry {
		b.Run(fmt.Sprintf("%04d", length), func(b *testing.B) {
			keys := make([]string, length)
			for k := 0; k < length; k++ {
				keys[k] = fmt.Sprintf("key%d", k)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m = make(map[string]int, length)
				for k := 0; k < length; k++ {
					m[keys[k]] = rand.Int()
				}
			}
			b.StopTimer()
			assert.Equal(b, length, len(m), "Ensure there are no duplicate keys")
			mapLiteral := builder()
			for k := range mapLiteral {
				_, ok := m[k]
				assert.Equal(b, true, ok, "Ensure %v is present", k)
			}
		})
	}
}

func BenchmarkMapStringLiteral(b *testing.B) {
	for length, builder := range buildersRegistry {
		b.Run(fmt.Sprintf("%04d", length), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				m = builder()
			}
		})
	}
}
