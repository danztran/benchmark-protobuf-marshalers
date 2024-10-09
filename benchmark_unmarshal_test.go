package just

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	gj "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func BenchmarkUnmarshal(b *testing.B) {
	benchData := genProduct(b)

	b.Run("Json", func(b *testing.B) {
		raw, err := json.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := json.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})

	b.Run("ProtoJson", func(b *testing.B) {
		raw, err := protojson.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := protojson.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})

	b.Run("Msgpack", func(b *testing.B) {
		raw, err := msgpack.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := msgpack.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})

	b.Run("Proto", func(b *testing.B) {
		raw, err := proto.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := proto.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})

	b.Run("Jsoniter", func(b *testing.B) {
		raw, err := jsoniter.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := jsoniter.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})

	b.Run("GoccyJson", func(b *testing.B) {
		raw, err := gj.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := gj.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})

	b.Run("Sonic", func(b *testing.B) {
		raw, err := sonic.Marshal(benchData)
		preBench(b, raw, err)
		for range b.N {
			prod := newProduct()
			err := sonic.Unmarshal(raw, prod)
			checkErr(b, err)
		}
	})
}

func preBench(b *testing.B, raw []byte, err error) {
	b.Helper()
	checkErr(b, err)
	b.SetBytes(int64(len(raw)))
	b.ResetTimer()
}
