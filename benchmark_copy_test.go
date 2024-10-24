package just

import (
	"encoding/json"
	"path"
	"slices"
	"testing"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
)

func BenchmarkCopy(b *testing.B) {
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
						x, err := engine.Marshal(item)
						if err != nil {
							b.Fatal(err)
						}
						prod := sample.New()
						if err := engine.Unmarshal(x, prod); err != nil {
							b.Fatal(err)
						}
					}
				})
			})
		}

		b.Run(path.Join(sample.Name, "Copier"), func(b *testing.B) {
			data, err := json.Marshal(item)
			if err != nil {
				b.Fatal(err)
			}
			b.RunParallel(func(bb *testing.PB) {
				b.SetBytes(int64(len(data)))
				b.ResetTimer()
				for bb.Next() {
					prod := sample.New()
					if err := copier.CopyWithOption(prod, item, copier.Option{
						DeepCopy: true,
					}); err != nil {
						b.Fatal(err)
					}
				}
			})
		})

		b.Run(path.Join(sample.Name, "ProtoClone"), func(b *testing.B) {
			data, err := json.Marshal(item)
			if err != nil {
				b.Fatal(err)
			}
			b.RunParallel(func(bb *testing.PB) {
				b.SetBytes(int64(len(data)))
				b.ResetTimer()
				for bb.Next() {
					_ = proto.Clone(item)
				}
			})
		})
	}
}
