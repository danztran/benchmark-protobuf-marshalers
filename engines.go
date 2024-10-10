package just

import (
	"encoding/json"

	"github.com/bytedance/sonic"
	gj "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Engine struct {
	Name      string
	Marshal   func(v any) ([]byte, error)
	Unmarshal func(data []byte, v any) error
}

type Marshaler func(v any) ([]byte, error)

type Unmarshaler func(data []byte, v any) error

var jsonEngine = Engine{
	Name:      "JsonBuiltin",
	Marshal:   json.Marshal,
	Unmarshal: json.Unmarshal,
}

var protojsonEngine = Engine{
	Name: "ProtoJson",
	Marshal: func(v any) ([]byte, error) {
		return protojson.Marshal(v.(proto.Message))
	},
	Unmarshal: func(data []byte, v any) error {
		return protojson.Unmarshal(data, v.(proto.Message))
	},
}

var protoEngine = Engine{
	Name: "Proto",
	Marshal: func(v any) ([]byte, error) {
		return proto.Marshal(v.(proto.Message))
	},
	Unmarshal: func(data []byte, v any) error {
		return proto.Unmarshal(data, v.(proto.Message))
	},
}

var msgpackEngine = Engine{
	Name:      "Msgpack",
	Marshal:   msgpack.Marshal,
	Unmarshal: msgpack.Unmarshal,
}

var jsoniterEngine = Engine{
	Name:      "Jsoniter",
	Marshal:   jsoniter.Marshal,
	Unmarshal: jsoniter.Unmarshal,
}

var goccyEngine = Engine{
	Name:      "GoccyDefault",
	Marshal:   gj.Marshal,
	Unmarshal: gj.Unmarshal,
}

var goccyUnorderedEngine = Engine{
	Name: "GoccyUnordered",
	Marshal: func(v any) ([]byte, error) {
		return gj.MarshalWithOption(v, gj.UnorderedMap())
	},
	Unmarshal: gj.Unmarshal,
}

var goccyFastestEngine = Engine{
	Name: "GoccyFastest",
	Marshal: func(v any) ([]byte, error) {
		return gj.MarshalWithOption(v, gj.UnorderedMap(), gj.DisableHTMLEscape(), gj.DisableNormalizeUTF8())
	},
	Unmarshal: gj.Unmarshal,
}

var (
	sonicDef    = sonic.ConfigDefault
	sonicEngine = Engine{
		Name:      "SonicDefault",
		Marshal:   sonicDef.Marshal,
		Unmarshal: sonicDef.Unmarshal,
	}
)

var (
	sonicCompact = sonic.Config{
		CompactMarshaler: true,
	}.Froze()
	sonicCompactEngine = Engine{
		Name:      "SonicCompact",
		Marshal:   sonicCompact.Marshal,
		Unmarshal: sonicCompact.Unmarshal,
	}
)

var (
	sonicStd = sonic.Config{
		EscapeHTML:       true,
		SortMapKeys:      false,
		CompactMarshaler: true,
		CopyString:       true,
		ValidateString:   true,
	}.Froze()
	sonicStdEngine = Engine{
		Name:      "SonicStd",
		Marshal:   sonicStd.Marshal,
		Unmarshal: sonicStd.Unmarshal,
	}
)

var (
	sonicFast = sonic.Config{
		NoQuoteTextMarshaler:    true,
		NoValidateJSONMarshaler: true,
		NoValidateJSONSkip:      true,
		CompactMarshaler:        true,
	}.Froze()
	sonicFastEngine = Engine{
		Name:      "SonicFast",
		Marshal:   sonicFast.Marshal,
		Unmarshal: sonicFast.Unmarshal,
	}
)

type Sample struct {
	Name   string
	New    func() proto.Message
	Random func() proto.Message
}
