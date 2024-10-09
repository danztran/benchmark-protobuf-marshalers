package just

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	gj "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	msgpack "github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func BenchmarkCopy(b *testing.B) {
	benchData := genProduct(b)

	b.Run("Json", func(b *testing.B) {
		for i := range b.N {
			raw, err := json.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = json.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("Proto", func(b *testing.B) {
		for i := range b.N {
			raw, err := proto.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = proto.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("ProtoJson", func(b *testing.B) {
		for i := range b.N {
			raw, err := protojson.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = protojson.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("Msgpack", func(b *testing.B) {
		for i := range b.N {
			raw, err := msgpack.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = msgpack.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("Jsoniter", func(b *testing.B) {
		for i := range b.N {
			raw, err := jsoniter.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = jsoniter.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("GoccyJson", func(b *testing.B) {
		for i := range b.N {
			raw, err := gj.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = gj.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("Sonic", func(b *testing.B) {
		for i := range b.N {
			raw, err := sonic.Marshal(benchData)
			checkErr(b, err)
			d := newProduct()
			err = sonic.Unmarshal(raw, d)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(raw)))
			}
		}
	})

	b.Run("ProtoClone", func(b *testing.B) {
		b.SetBytes(int64(proto.Size(benchData)))
		b.ResetTimer()
		for range b.N {
			_ = protoClone(benchData)
			// _ = d
		}
	})
}
