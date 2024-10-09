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

var (
	jsonData, _      = json.Marshal(benchData)
	protoData, _     = proto.Marshal(benchData)
	protoJsonData, _ = protojson.Marshal(benchData)
	jsoniterData, _  = jsoniter.Marshal(benchData)
	goccyJsonData, _ = gj.Marshal(benchData)
	msgpackData, _   = msgpack.Marshal(benchData)
	sonicData, _     = sonic.Marshal(benchData)
)

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("Json", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = json.Unmarshal(jsonData, data)
		}
	})

	b.Run("Proto", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = proto.Unmarshal(protoData, data)
		}
	})

	b.Run("ProtoJson", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = protojson.Unmarshal(protoJsonData, data)
		}
	})

	b.Run("Jsoniter", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = jsoniter.Unmarshal(jsoniterData, data)
		}
	})

	b.Run("GoccyJson", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = gj.Unmarshal(goccyJsonData, data)
		}
	})

	b.Run("Msgpack", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = msgpack.Unmarshal(msgpackData, data)
		}
	})

	b.Run("Sonic", func(b *testing.B) {
		for range b.N {
			data := newProduct()
			_ = sonic.ConfigStd.Unmarshal(sonicData, data)
		}
	})
}
