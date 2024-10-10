package just

import (
	"path"
	"slices"
	"testing"
)

func BenchmarkMarshal(b *testing.B) {
	for sample := range slices.Values(samples) {
		item := sample.Random()
		for engine := range slices.Values(engines) {
			data, err := engine.Marshal(item)
			if err != nil {
				b.Fatal(err)
			}
			b.Run(path.Join(sample.Name, engine.Name), func(b *testing.B) {
				b.SetBytes(int64(len(data)))
				b.ResetTimer()
				b.RunParallel(func(bb *testing.PB) {
					for bb.Next() {
						_, err := engine.Marshal(item)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})
		}
	}
}
