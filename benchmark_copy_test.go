package just

import (
	"encoding/json"
	"testing"

	gj "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	msgpack "github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func BenchmarkJsonCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := json.Marshal(benchData)
		d := newProduct()
		_ = json.Unmarshal(raw, d)
	}
}

func BenchmarkProtoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := proto.Marshal(benchData)
		d := newProduct()
		_ = proto.Unmarshal(raw, d)
	}
}

func BenchmarkProtoJsonCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := protojson.Marshal(benchData)
		d := newProduct()
		_ = protojson.Unmarshal(raw, d)
	}
}

func BenchmarkMsgpackCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := msgpack.Marshal(benchData)
		d := newProduct()
		_ = msgpack.Unmarshal(raw, d)
	}
}

func BenchmarkJsoniterCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := jsoniter.Marshal(benchData)
		d := newProduct()
		_ = jsoniter.Unmarshal(raw, d)
	}
}

func BenchmarkGoccyJsonCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		raw, _ := gj.Marshal(benchData)
		d := newProduct()
		_ = gj.Unmarshal(raw, d)
	}
}

func BenchmarkProtoCloneCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := protoClone(benchData)
		_ = d
	}
}
