package just

import (
	"testing"

	"github.com/goccy/go-json"
	gj "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestJson(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	raw, err := json.Marshal(data)
	require.NoError(t, err)

	d := newProduct()
	err = json.Unmarshal(raw, d)
	require.NoError(t, err)

	equal := proto.Equal(data, d)
	require.True(t, equal)
}

func TestGrpc(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	raw, err := proto.Marshal(data)
	require.NoError(t, err)

	d := newProduct()
	err = proto.Unmarshal(raw, d)
	require.NoError(t, err)

	equal := proto.Equal(data, d)
	require.True(t, equal)
}

func TestProtoClone(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	d := protoClone(data)
	equal := proto.Equal(data, d)
	require.True(t, equal)
}

func TestProtoJson(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	raw, err := protojson.Marshal(data)
	require.NoError(t, err)

	d := newProduct()
	err = protojson.Unmarshal(raw, d)
	require.NoError(t, err)

	equal := proto.Equal(data, d)
	require.True(t, equal)
}

func TestJsoniter(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	raw, err := jsoniter.Marshal(data)
	require.NoError(t, err)

	d := newProduct()
	err = jsoniter.Unmarshal(raw, d)
	require.NoError(t, err)

	equal := proto.Equal(data, d)
	require.True(t, equal)
}

func TestGoccyJson(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	raw, err := gj.Marshal(data)
	require.NoError(t, err)

	d := newProduct()
	err = gj.Unmarshal(raw, d)
	require.NoError(t, err)

	equal := proto.Equal(data, d)
	require.True(t, equal)
}

func TestMsgpack(t *testing.T) {
	t.Parallel()
	data := genProduct(t)
	require.NotZero(t, data)

	raw, err := msgpack.Marshal(data)
	require.NoError(t, err)

	d := newProduct()
	err = msgpack.Unmarshal(raw, d)
	require.NoError(t, err)

	equal := proto.Equal(data, d)
	require.True(t, equal)
}
