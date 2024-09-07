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

func BenchmarkJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(benchData)
	}
}

func BenchmarkProtoMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(benchData)
	}
}

func BenchmarkProtoJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = protojson.Marshal(benchData)
	}
}

func BenchmarkJsoniterMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = jsoniter.Marshal(benchData)
	}
}

func BenchmarkGoccyJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = gj.Marshal(benchData)
	}
}

func BenchmarkMsgpackMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(benchData)
	}
}
