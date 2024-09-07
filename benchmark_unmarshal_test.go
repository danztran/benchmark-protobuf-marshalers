package just

import (
	"encoding/json"
	"testing"

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
)

func BenchmarkJsonUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := newProduct()
		_ = json.Unmarshal(jsonData, data)
	}
}

func BenchmarkProtoUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := newProduct()
		_ = proto.Unmarshal(protoData, data)
	}
}

func BenchmarkProtoJsonUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := newProduct()
		_ = protojson.Unmarshal(protoJsonData, data)
	}
}

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := newProduct()
		_ = jsoniter.Unmarshal(jsoniterData, data)
	}
}

func BenchmarkGoccyJsonUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := newProduct()
		_ = gj.Unmarshal(goccyJsonData, data)
	}
}

func BenchmarkMsgpackUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := newProduct()
		_ = msgpack.Unmarshal(msgpackData, data)
	}
}
