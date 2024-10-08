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
	b.Run("JsonCopy", func(b *testing.B) {
		for range b.N {
			raw, _ := json.Marshal(benchData)
			d := newProduct()
			_ = json.Unmarshal(raw, d)
		}
	})

	b.Run("ProtoCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			raw, _ := proto.Marshal(benchData)
			d := newProduct()
			_ = proto.Unmarshal(raw, d)
		}
	})

	b.Run("ProtoJsonCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			raw, _ := protojson.Marshal(benchData)
			d := newProduct()
			_ = protojson.Unmarshal(raw, d)
		}
	})

	b.Run("MsgpackCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			raw, _ := msgpack.Marshal(benchData)
			d := newProduct()
			_ = msgpack.Unmarshal(raw, d)
		}
	})

	b.Run("JsoniterCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			raw, _ := jsoniter.Marshal(benchData)
			d := newProduct()
			_ = jsoniter.Unmarshal(raw, d)
		}
	})

	b.Run("GoccyJsonCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			raw, _ := gj.Marshal(benchData)
			d := newProduct()
			_ = gj.Unmarshal(raw, d)
		}
	})

	b.Run("SonicCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			raw, _ := sonic.ConfigStd.Marshal(benchData)
			d := newProduct()
			_ = sonic.ConfigStd.Unmarshal(raw, d)
		}
	})

	b.Run("ProtoCloneCopy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := protoClone(benchData)
			_ = d
		}
	})
}
