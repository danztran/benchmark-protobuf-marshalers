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
	b.Run("Json", func(b *testing.B) {
		for range b.N {
			_, err := json.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Proto", func(b *testing.B) {
		for range b.N {
			_, err := proto.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("ProtoJson", func(b *testing.B) {
		for range b.N {
			_, err := protojson.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Jsoniter", func(b *testing.B) {
		for range b.N {
			_, err := jsoniter.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("GoccyJson", func(b *testing.B) {
		for range b.N {
			_, err := gj.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Msgpack", func(b *testing.B) {
		for range b.N {
			_, err := msgpack.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Sonic", func(b *testing.B) {
		for range b.N {
			_, err := sonic.Marshal(benchData)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
