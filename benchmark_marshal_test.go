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

func BenchmarkMarshal(b *testing.B) {
	benchData := genProduct(b)

	b.Run("Json", func(b *testing.B) {
		for i := range b.N {
			data, err := json.Marshal(benchData)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(data)))
			}
		}
	})

	b.Run("ProtoJson", func(b *testing.B) {
		for i := range b.N {
			data, err := protojson.Marshal(benchData)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(data)))
			}
		}
	})

	b.Run("Msgpack", func(b *testing.B) {
		for i := range b.N {
			data, err := msgpack.Marshal(benchData)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(data)))
			}
		}
	})

	b.Run("Proto", func(b *testing.B) {
		for i := range b.N {
			data, err := proto.Marshal(benchData)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(data)))
			}
		}
	})

	b.Run("Jsoniter", func(b *testing.B) {
		for i := range b.N {
			data, err := jsoniter.Marshal(benchData)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(data)))
			}
		}
	})

	b.Run("GoccyJson", func(b *testing.B) {
		for i := range b.N {
			data, err := gj.Marshal(benchData)
			checkErr(b, err)
			if i == 0 {
				b.SetBytes(int64(len(data)))
			}
		}
	})

	b.Run("Sonic", func(b *testing.B) {
		for range b.N {
			_, err := sonic.Marshal(benchData)
			checkErr(b, err)
		}
	})
}
